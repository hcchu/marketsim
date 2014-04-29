package main

import (
    "fmt"
    "os"
    "bufio"

    "github.com/hcchu/marketsim/marketsim"
)

// "2014-04-01T16:19:00Z BUY USD 100.000 100"

func main() {
	reader := bufio.NewReader(os.Stdin)
    bid_book := marketsim.NewBook()
    ask_book := marketsim.NewBook()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
            fmt.Println(err)
			break
		}
        orderline := marketsim.ParseOrder(line)
        marketsim.DispatchOrder(orderline, &bid_book, &ask_book)
        //order_book = append(order_book, *orderline)
        //sort.Sort(marketsim.ByTimestamp(order_book))
        //sort.Sort(marketsim.ByPrice(order_book))
        //fmt.Println(order_book)
        //fmt.Println(order_book[0])
        fmt.Println(bid_book)
        fmt.Println(ask_book)
	}
}
