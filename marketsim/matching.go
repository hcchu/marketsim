package marketsim

import "fmt"
import "sort"

func removeOrder(b []Order) {
    b = append(b[:0], b[1:]...)
}

func MatchOrder(o *Order, b *map[OrderKey][]Order) bool {
    sb := (*b)[OrderKey{o.Commodity, "SELL"}]
    bb := (*b)[OrderKey{o.Commodity, "BUY"}]
    switch {
    //case o.OrderType == "BUY" && len((*b)[OrderKey{o.Commodity, "SELL"}]) == 0:
    case o.OrderType == "BUY" && len(sb) == 0:
        return false
    case o.OrderType == "SELL" && len(bb) == 0:
        return false
    case o.OrderType == "BUY":
        if o.Price >= sb[0].Price {
            switch {
            case o.Amount > sb[0].Amount:
                fmt.Println("BUY ORDER EXECUTED for", sb[0].Amount, o.Commodity, "at", o.Price)
                (*o).Amount = o.Amount - sb[0].Amount
                //(*b)[OrderKey{o.Commodity, o.OrderType}] = append((*b)[OrderKey{o.Commodity, o.OrderType}], *o)
                (*b)[OrderKey{o.Commodity, o.OrderType}] = append(bb, *o)
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
        if o.Price <= bb[0].Price {
            fmt.Println("SELL ORDER EXECUTED for",
            o.Commodity, "at", o.Price)
            removeOrder((*b)[OrderKey{o.Commodity, "BUY"}])
            return true
        }
        return false
    }
    return false
}




