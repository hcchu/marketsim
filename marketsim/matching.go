package marketsim

import "fmt"

func MatchOrder(o *Order, b *map[OrderKey][]Order) bool {
	sb := (*b)[OrderKey{o.Commodity, "SELL"}]
	bb := (*b)[OrderKey{o.Commodity, "BUY"}]
	switch {
	case o.OrderType == "BUY" && len(sb) == 0:
		return false
	case o.OrderType == "SELL" && len(bb) == 0:
		return false
	case o.OrderType == "BUY":
		if o.Price >= sb[0].Price {
			switch {
			case o.Amount > sb[0].Amount:
				fmt.Println("BUY ORDER", o.OrderID, "EXECUTED for", sb[0].Amount, o.Commodity, "at", o.Price, "matching", sb[0].OrderID)
				(*o).Amount = o.Amount - sb[0].Amount
				removeOrder(b, o.Commodity, "SELL")
				// attempt to fill rest of order
				DispatchOrder(o, b)
				return true
			case o.Amount < sb[0].Amount:
				fmt.Println("BUY ORDER", o.OrderID, "EXECUTED for", o.Amount, o.Commodity, "at", o.Price, "matching", sb[0].OrderID)
				(*b)[OrderKey{o.Commodity, "SELL"}][0].Amount = sb[0].Amount - o.Amount
				return true
			case o.Amount == sb[0].Amount:
				fmt.Println("BUY ORDER", o.OrderID, "EXECUTED for", o.Amount, o.Commodity, "at", o.Price, "matching", sb[0].OrderID)
				removeOrder(b, o.Commodity, "SELL")
				return true
			}
		}
	case o.OrderType == "SELL":
		if o.Price <= bb[0].Price {
			switch {
			case o.Amount > bb[0].Amount:
				fmt.Println("SELL ORDER", o.OrderID, "EXECUTED for", bb[0].Amount, o.Commodity, "at", o.Price, "matching", bb[0].OrderID)
				(*o).Amount = o.Amount - bb[0].Amount
				removeOrder(b, o.Commodity, "BUY")
				// attempt to fill rest of order
				DispatchOrder(o, b)
				return true
			case o.Amount < bb[0].Amount:
				fmt.Println("SELL ORDER", o.OrderID, "EXECUTED for", o.Amount, o.Commodity, "at", o.Price, "matching", bb[0].OrderID)
				(*b)[OrderKey{o.Commodity, "BUY"}][0].Amount = bb[0].Amount - o.Amount
				return true
			case o.Amount == bb[0].Amount:
				fmt.Println("SELL ORDER", o.OrderID, "EXECUTED for", o.Amount, o.Commodity, "at", o.Price, "matching", bb[0].OrderID)
				removeOrder(b, o.Commodity, "BUY")
				return true
			}
		}
		return false
	}
	return false
}
