package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hcchu/marketsim/marketsim"
)

// "2014-04-01T16:19:00Z BUY USD 100.000 100"

func main() {
	reader := bufio.NewReader(os.Stdin)

	order_book := marketsim.NewOrderBook()

    order_id := 1000

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		orderline := marketsim.ParseOrder(line, order_id)
		marketsim.DispatchOrder(orderline, &order_book)
        order_id += 1
	}
}
