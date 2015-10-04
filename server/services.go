package server

import ("strings"
        "fmt"
        "io/ioutil"
        "encoding/json"
        "net/http"
        "strconv"
)

/*
  This structure will be used by JSON data returned from yahoo finance API
*/
type Jsonresponse struct {
  List struct {
    Resources []struct {
      Resource struct {
          Fields struct {
            Name string
            Price string
            Symbol string
          }
      }
    }
  }
}

var noOfShares map[string]int // This map will save share as key, and no of shares purchased by the user as Values
var url = "http://finance.yahoo.com/webservice/v1/symbols/"

/*
  This function will create Yahoo finance API URL from user input
*/
func makeUrl(stocksWithPercentage []string) (map[string]string, string){
  stocksData := make(map[string]string)
  var symbolStr string
   for data := range stocksWithPercentage {
       singleStock := strings.Split(stocksWithPercentage[data],":")
       percentage := strings.Split(singleStock[1],"%")
       stocksData[singleStock[0]] = percentage[0]
  }

  for symbol := range stocksData {
 		symbolStr += symbol + ","
 	}
 	symbolStr = symbolStr[:len(symbolStr)-1]
  urlCreated := url+symbolStr+"/quote?format=json&view=‌detail"
  fmt.Println("url: ",urlCreated)
  return stocksData, urlCreated
}

/*
  This function will call get on yahoo finance api to fetch JSON data, and populate Jsonresponse structure
*/
func DataFromYahoo(url string) {
  resp, err := http.Get(url)
  if err!=nil {
    fmt.Println("Error from DataFromYahoo function: ",err)
  } else {
  defer resp.Body.Close()
         jsonDataFromHttp, err := ioutil.ReadAll(resp.Body)
         if err != nil {
                 panic(err)
         }
         err = json.Unmarshal([]byte(jsonDataFromHttp), &DataPopulated)

         if err != nil {
           fmt.Println("error: ",err)
                 panic(err)
         }
   }
}

/*
This map will have mapping of shares requestes by the user to buy, to its current market price.
*/
func makeStockWithCurrentPriceMap(data Jsonresponse) map[string]float64{
  stocksPrice := make(map[string]float64)
   for _, resourceList := range DataPopulated.List.Resources {
     price, err := strconv.ParseFloat(resourceList.Resource.Fields.Price,64)
     if err!= nil {
       fmt.Println(err)
     } else {
    stocksPrice[resourceList.Resource.Fields.Symbol] = price
     }
   }
return stocksPrice
}

/*
  This function will calculate budget for each share requested by the user,
  no of shares bought in that budget allocated for the particular share
  uninvestedAmount that remains with the user.
  Finally, this function will populate Reply structure that will be returned to the user
*/
func stocksPurchased(budget float32,stocksData map[string]string, stockPrice map[string]float64, reply *Reply)  {
   var uninvestedAmount float64
   uninvestedAmount = float64(budget)
   noOfShares = make(map[string]int)
  for share, percentage := range stocksData {
     percentage, err := strconv.ParseFloat(percentage,64)
     if err == nil {
       budgetForShare := (float64(budget) * (percentage/100))
       uninvestedAmount = float64(uninvestedAmount) - budgetForShare
       sharesCount := budgetForShare /stockPrice[share]
       noOfShares[share] = int(sharesCount)
       uninvestedAmount += budgetForShare - float64(int(sharesCount))*stockPrice[share]
      }
  }

var stocks string
  for share, count := range noOfShares {
   stocks+=string(share)+":"+ strconv.Itoa(count) +":$" + strconv.FormatFloat(stockPrice[share], 'G', -1, 64) + ","
  }
  stocks = stocks[:len(stocks)-1]
  reply.UnivestedAmount = float32(uninvestedAmount)
  reply.Stocks = stocks
  tradeId = tradeId+1
  reply.TradeId = tradeId
  tradeIdWithSharesMap[reply.TradeId] = noOfShares
  tradeIdWithUninvestedAmount[reply.TradeId] = uninvestedAmount
}

/*
This function will populate PortFolioReply strucure that will be returend to the user, when user request portfolio view request.
*/
func FetchPortFolioRecords(tradeId int,reply *PortFolioReply) {
  var symbolStr string
  var stocks string
  var marketValue float64

 if len(tradeIdWithSharesMap[tradeId]) == 0 {
  reply.ErrorMessage = "No portfolio available for TradeId "+strconv.Itoa(tradeId)+". Please buy some shares first"
  return
}
  sharesWithNo :=  tradeIdWithSharesMap[tradeId]
  for share, _ := range sharesWithNo {
    symbolStr +=share+","
  }
 	symbolStr = symbolStr[:len(symbolStr)-1]
  urlCreated := url+symbolStr+"/quote?format=json&view=‌detail"
  DataFromYahoo(urlCreated)
  stocksPrice := makeStockWithCurrentPriceMap(DataPopulated)
  for share, count := range sharesWithNo {
    stocks += share+":"+  strconv.Itoa(count)+":$"+ strconv.FormatFloat(stocksPrice[share], 'G', -1, 64)+","
    marketValue += stocksPrice[share]*float64(count)
  }
 reply.Stocks = stocks[:len(stocks)-1]
 reply.CurrentMarketValue = float32(marketValue)
 reply.UnivestedAmount = float32(tradeIdWithUninvestedAmount[tradeId])
}
