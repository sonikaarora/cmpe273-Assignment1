package server

import (
      "strings"
)

var DataPopulated Jsonresponse

func(this *ServerStruct) BuyingStocks(request Request, reply *Reply) error{

   stocksWithPercentage := strings.Split(request.StockSymbolAndPercentage, ",")
   stocksPercentageMap, url := makeUrl(stocksWithPercentage)
   DataFromYahoo(url)
   stockPrieMap := makeStockWithCurrentPriceMap(DataPopulated)
   stocksPurchased(request.Budget,stocksPercentageMap,stockPrieMap,reply)

   return nil
}

func (this *PortfolioStruct) PortfolioView(tradeId int, reply *PortFolioReply) error {
  FetchPortFolioRecords(tradeId,reply)
  return nil
}
