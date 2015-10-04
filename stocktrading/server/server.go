package server

import (
  "fmt"
  "net/rpc"
  "net/http"
)
type ServerStruct struct {}
type PortfolioStruct struct {}


type Request struct {
    StockSymbolAndPercentage string
    Budget float32
}

type Reply struct {
    TradeId int
    Stocks string
    UnivestedAmount float32
}

type PortFolioReply struct {
    Stocks string
    CurrentMarketValue float32
    UnivestedAmount float32
    ErrorMessage string
}
var tradeId int = 0 // This is tradeID, and will be incremented by one for next user. This will work only for 1 session.
var tradeIdWithSharesMap map[int]map[string]int //This map will have mapping of trade id with noOfShares map
var tradeIdWithUninvestedAmount map[int]float64


func Init() { // These maps will be queried to view portfolio
  tradeIdWithSharesMap = make(map[int]map[string]int)
  tradeIdWithUninvestedAmount = make(map[int]float64)
}

func Server() {
  Init()
  rpc.Register(new(ServerStruct))
  rpc.Register(new(PortfolioStruct))
  rpc.HandleHTTP()
  fmt.Println("Server listining on 9999 port.....")
  err := http.ListenAndServe(":9999",nil)
  if err != nil {
    fmt.Println(err)
    return
  }
}
