package marketsim

import "fmt"
import "sort"

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
            switch {
            case o.Amount > ((*b)[OrderKey{o.Commodity, "SELL"}][0].Amount):
                fmt.Println("BUY ORDER EXECUTED for", ((*b)[OrderKey{o.Commodity, "SELL"}][0].Amount), o.Commodity, "at", o.Price)
                (*o).Amount = o.Amount - ((*b)[OrderKey{o.Commodity, "SELL"}][0].Amount)
                (*b)[OrderKey{o.Commodity, o.OrderType}] = append((*b)[OrderKey{o.Commodity, o.OrderType}], *o)
                sort.Sort(ByTimestamp((*b)[OrderKey{o.Commodity, o.OrderType}]))
                sort.Sort(ByPrice((*b)[OrderKey{o.Commodity, o.OrderType}]))
                removeOrder((*b)[OrderKey{o.Commodity, "SELL"}])
                fmt.Println(*b)
            }

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




