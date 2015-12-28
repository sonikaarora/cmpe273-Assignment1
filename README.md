# cmpe273-Assignment1

----------------------
Design

Server.go: 
-    the server will listen on port 9999
-    the server will maintain three structures,
Request: Client request structure
Reply: Buy stock reply structure
PortFolioReply: Portfolio reply structure
-    Init function will initialize maps that will be queried against to serve portfolio request from user
Controller.go 
-    BuyingStocks function is called when, user initiates request to buy stocks
-    PortfolioView function is called when, user initiates request to view portfolio
Services.go
-    Jsonresponse structure will be used by JSON data returned from yahoo finance API
-    DataFromYahoo function will fetch stocks data from yahoo finance API.
-    The stockspurchased function will do all the calculations and populate reply structure for buying stocks request from the user.
-    FetchPortFolioRecords function will fetch data from the local cache and populate PortFolioReply structure.

Client.go: Client will initiate HTTP connection with server



----------------------
Command Line Arguments

1) Start server using the following command, server will start listening on port 9999
    Go run startserver.go

2) Start client and provide appropriate command line arguments to buy stocks or to view portfolio

-  Buy Stocks
    go run startclient.go buy stock1:percentage,stock:percentage budget

-  View Portfolio
    go run startclient.go view tradeId  




