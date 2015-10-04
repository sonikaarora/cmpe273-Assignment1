package main
import (
  "fmt"
  "stocktrading/server"
)

func main() {
 go server.Server()

 var input int
 fmt.Scanln(&input)


}
