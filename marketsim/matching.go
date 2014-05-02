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

func removeOrder(book *[]Order) {
    copy((*book)[0:], (*book)[1:])
    (*book)[len(*book)-1] = Order{}
    *book = (*book)[:len(*book)-1]
}

func MatchOrder(o *Order, sell *[]Order, buy *[]Order) bool {
    switch {
    case o.OrderType == "BUY" && len(*sell) == 0: 
        return false
    case o.OrderType == "SELL" && len(*buy) == 0: 
        return false
    case o.OrderType == "BUY":
        if o.Price >= (*sell)[0].Price {
            fmt.Println("BUY ORDER EXECUTED for",
            o.Commodity, "at", o.Price)
            removeOrder(sell)
        }
        return true
    case o.OrderType == "SELL":
        if o.Price <= (*buy)[0].Price {
            fmt.Println("SELL ORDER EXECUTED for",
            o.Commodity, "at", o.Price)
            removeOrder(buy)
        }
        return true
    }
    return false
}

