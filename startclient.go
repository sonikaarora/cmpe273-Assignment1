package main
import (
  "fmt"
  "stocktrading/client"
)

func main() {
 go client.Client()

 var input int
 fmt.Scanln(&input)


}
