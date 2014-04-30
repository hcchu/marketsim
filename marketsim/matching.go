package marketsim

import "fmt"

func MatchEngine(b []Order, s []Order) {
    switch {
    case len(s) == 0:
        fmt.Println("Sell order book empty")
    case len(b) == 0:
        fmt.Println("Buy order book empty")
    case b[0].Price == s[0].Price:
        fmt.Println("TRADE EXECUTED")
    }
}
