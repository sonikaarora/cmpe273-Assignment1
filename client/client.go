package client

import (
  "fmt"
  "net/rpc"
  "os"
  "strconv"
)

/*
Request structure that would be populated with client request parameters, when user request buying of stocks.
*/
type Request struct {
    StockSymbolAndPercentage string
    Budget float32
}

/*
Reply structure returned to the user, when request to buy stock is served
*/
type Reply struct {
    TradeId int
    Stocks string
    UnivestedAmount float32
}
/*
Portfolio Reply structure returned to the user, when request to view portfolio is served
*/

type PortFolioReply struct {
    Stocks string
    CurrentMarketValue float32
    UnivestedAmount float32
    ErrorMessage string
}

func Client() {
  c, err := rpc.DialHTTP("tcp","127.0.0.1:9999")
  if err != nil {
    fmt.Println(err)
    return
  }

var reply Reply
var portfolioReply PortFolioReply
var Shares Request
var TradiIdRequested int

// Error Handling
if os.Args == nil || len(os.Args) < 3 {
    fmt.Println("Please enter valid command line arguments e.g. Buy GOOG:50%,YHOO:50% 1000")
    return
  }

if os.Args == nil || len(os.Args) < 4 {
   if os.Args[1] == "buy" {
     fmt.Println("Please enter valid command line arguments e.g. Buy GOOG:50%,YHOO:50% 1000")
     return
   }
}

userAction := os.Args[1]
if userAction == "buy" {
  stockString := os.Args[2]
  budgetInFloat64, _ := strconv.ParseFloat(os.Args[3], 64)
  budget := float32(budgetInFloat64)
  Shares = Request{StockSymbolAndPercentage:stockString, Budget:budget}

}
if userAction == "view" {
 TradiIdRequested,_ = strconv.Atoi(os.Args[2])
}

switch userAction {
case "buy":
  err = c.Call("ServerStruct.BuyingStocks", Shares, &reply) // Synchronous call

   if err!= nil {
     fmt.Println(err)
   } else {
  fmt.Println("reply from buying stock:\ntradeId: ", reply.TradeId," \nStocks: ",reply.Stocks,"\nUnivestedAmount: ",reply.UnivestedAmount )
   }
case "view":
  err = c.Call("PortfolioStruct.PortfolioView",TradiIdRequested, &portfolioReply) // Synchronous call
   if err!= nil {
     fmt.Println(err)
   }
   if len(portfolioReply.ErrorMessage) > 0{
     fmt.Println(portfolioReply.ErrorMessage)
     } else {
       fmt.Println("Reply from view:\nstocks: ",portfolioReply.Stocks,"\nCurrentMarketValue: ",portfolioReply.CurrentMarketValue,"\nUnivestedAmount: ",portfolioReply.UnivestedAmount)
   }
 }

}
