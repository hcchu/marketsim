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

func removeOrder(b []Order) {
    b = append(b[:0], b[1:]...)
}

func MatchOrder(o *Order, b *map[OrderKey][]Order) bool {
    switch {
    case o.OrderType == "BUY" && len((*b)[OrderKey{o.Commodity, "SELL"}]) == 0:
        return false
    case o.OrderType == "SELL" && len((*b)[OrderKey{o.Commodity, "BUY"}]) == 0:
        return false
    case o.OrderType == "BUY":
        if o.Price >= ((*b)[OrderKey{o.Commodity, "SELL"}][0].Price) {
            fmt.Println("BUY ORDER EXECUTED for", o.Commodity, "at", o.Price)
            removeOrder((*b)[OrderKey{o.Commodity, "SELL"}])
            return true
        }
        return false
    case o.OrderType == "SELL":
        if o.Price <= ((*b)[OrderKey{o.Commodity, "BUY"}][0].Price) {
            fmt.Println("SELL ORDER EXECUTED for",
            o.Commodity, "at", o.Price)
            removeOrder((*b)[OrderKey{o.Commodity, "BUY"}])
            return true
        }
        return false
    }
    return false
}




