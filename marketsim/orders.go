package marketsim

import (
    "strings"
    "strconv"
    "time"
    "sort"
)

type Order struct {
	Timestamp  int64
	OrderType string
	Commodity  string
	Price      float32
	Amount     int32
}

type ByTimestamp []Order

func (a ByTimestamp) Len() int { return len(a) }
func (a ByTimestamp) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTimestamp) Less(i, j int) bool { return a[i].Timestamp > a[j].Timestamp }

type ByPrice []Order

func (a ByPrice) Len() int { return len(a) }
func (a ByPrice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
// Sort bid order by highest price, ask order by lowest 
func (a ByPrice) Less(i, j int) bool { 
    if a[i].OrderType == "BUY" {
        return a[i].Price > a[j].Price 
    }
    return a[i].Price < a[j].Price
}

func NewBook() []Order {
    o := make([]Order, 0)
    return o
}
        
// Reads a string from stdin and returns Order type
func ParseOrder(o string) *Order {
    order_string := strings.Fields(o)
    p := new(Order)
    timestamp, _ := time.Parse(time.RFC3339, order_string[0])
    p.Timestamp = timestamp.Unix()
    p.OrderType = strings.ToUpper(order_string[1])
    p.Commodity = order_string[2]
    price, _ := strconv.ParseFloat(order_string[3], 32)
    p.Price = float32(price)
    amount, _ := strconv.ParseInt(order_string[4], 10, 32)
    p.Amount = int32(amount)
	return p
}

// Send order to the proper order book
func DispatchOrder(o *Order, b *[]Order, a *[]Order) {
    switch {
    case o.OrderType == "BUY":
        *b = append(*b, *o)
        sort.Sort(ByTimestamp(*b))
        sort.Sort(ByPrice(*b))
    case o.OrderType == "SELL":
        *a = append(*a, *o)
        sort.Sort(ByTimestamp(*a))
        sort.Sort(ByPrice(*a))
    }
}

