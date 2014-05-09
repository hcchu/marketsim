package marketsim

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

type Order struct {
	Timestamp int64
	OrderType string
	Price     float32
	Commodity string
	Amount    int32
}

type OrderKey struct {
	Commodity, OrderType string
}

type ByTimestamp []Order

func (a ByTimestamp) Len() int           { return len(a) }
func (a ByTimestamp) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTimestamp) Less(i, j int) bool { return a[i].Timestamp < a[j].Timestamp }

type ByPrice []Order

func (a ByPrice) Len() int      { return len(a) }
func (a ByPrice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Sort buy order by highest price, sell order by lowest
func (a ByPrice) Less(i, j int) bool {
	if a[i].OrderType == "BUY" {
		return a[i].Price > a[j].Price
	}
	return a[i].Price < a[j].Price
}

func NewOrderBook() map[OrderKey][]Order {
	o := make(map[OrderKey][]Order)
	return o
}

// Reads a string from stdin and returns Order type
func ParseOrder(o string) *Order {
	order_string := strings.Fields(o)
	p := new(Order)
	timestamp, _ := time.Parse(time.RFC3339, order_string[0])
	p.Timestamp = timestamp.Unix()
	p.OrderType = strings.ToUpper(order_string[1])
	price, _ := strconv.ParseFloat(order_string[2], 32)
	p.Price = float32(price)
	p.Commodity = order_string[3]
	amount, _ := strconv.ParseInt(order_string[4], 10, 32)
	p.Amount = int32(amount)
	return p
}

// Send order to the proper order book
func DispatchOrder(o *Order, b *map[OrderKey][]Order) {
	result := MatchOrder(o, b)
	if result == false {
		(*b)[OrderKey{o.Commodity, o.OrderType}] = append((*b)[OrderKey{o.Commodity, o.OrderType}], *o)
		sort.Sort(ByTimestamp((*b)[OrderKey{o.Commodity, o.OrderType}]))
		sort.Sort(ByPrice((*b)[OrderKey{o.Commodity, o.OrderType}]))
	}
}

/*
func removeOrder(b *[]Order) {
    b = append(b[:0], b[1:]...)
}
*/

func removeOrder(b *map[OrderKey][]Order, c string, s string) {
	(*b)[OrderKey{c, s}] = append((*b)[OrderKey{c, s}][:0],
		(*b)[OrderKey{c, s}][1:]...)
}
