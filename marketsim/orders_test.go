package marketsim

import "testing"
import "sort"

func TestParseOrder(t *testing.T) {
	in := "2014-04-01T16:19:00Z BUY 100.000 USD 100"
	out := Order{1396369140, "BUY", 100, "USD", 100, 1000}
	if x := ParseOrder(in, 1000); *x != out {
		t.Errorf("ParseOrder(%v) = %v, want %v", in, x, out)
	}
}

func TestDispatchOrder(t *testing.T) {
	order_in := Order{1396369140, "BUY", 100, "USD", 100, 1000}
	book_in := make(map[OrderKey][]Order)
	book_out := make(map[OrderKey][]Order)
	book_out[OrderKey{"USD", "BUY"}] = []Order{{1396369140, "BUY", 100, "USD", 100, 1000}}
	DispatchOrder(&order_in, &book_in)
	if book_in[OrderKey{"USD", "BUY"}][0] != book_out[OrderKey{"USD", "BUY"}][0] {
		t.Errorf("DispatchOrder = %v, want %v", book_in, book_out)
	}
}

func TestBuySort(t *testing.T) {
	buy_book := []Order{{1396369141, "BUY", 100, "USD", 100, 1000}, {1396369141, "BUY", 101, "USD", 100, 1001}, {1396369140, "BUY", 101, "USD", 100, 1002}}
	buy_book_sorted := []Order{{1396369140, "BUY", 101, "USD", 100, 1002}, {1396369141, "BUY", 101, "USD", 100, 1001}, {1396369141, "BUY", 100, "USD", 100, 1000}}
	sort.Sort(ByTimestamp(buy_book))
	sort.Sort(ByPrice(buy_book))
	for index, element := range buy_book {
		if element != buy_book_sorted[index] {
			t.Errorf("Buy book = %v, want %v", buy_book, buy_book_sorted)
		}
	}
}

func TestSellSort(t *testing.T) {
	sell_book := []Order{{1396369141, "SELL", 100, "USD", 100, 1000}, {1396369141, "SELL", 101, "USD", 100, 1001}, {1396369140, "SELL", 101, "USD", 100, 1002}}
	sell_book_sorted := []Order{{1396369141, "SELL", 100, "USD", 100, 1000}, {1396369140, "SELL", 101, "USD", 100, 1002}, {1396369141, "SELL", 101, "USD", 100, 1001}}
	sort.Sort(ByTimestamp(sell_book))
	sort.Sort(ByPrice(sell_book))
	for index, element := range sell_book {
		if element != sell_book_sorted[index] {
			t.Errorf("Sell book = %v, want %v", sell_book, sell_book_sorted)
		}
	}
}

func TestPartialBuy(t *testing.T) {
	order_in := Order{1396369141, "BUY", 100, "USD", 100, 1000}
	book_in := make(map[OrderKey][]Order)
	book_out := make(map[OrderKey][]Order)
	book_in[OrderKey{"USD", "SELL"}] = []Order{{1396369140, "SELL", 100, "USD", 60, 1001}}
	book_out[OrderKey{"USD", "BUY"}] = []Order{{1396369141, "BUY", 100, "USD", 40, 1000}}
	DispatchOrder(&order_in, &book_in)
	if book_in[OrderKey{"USD", "BUY"}][0] != book_out[OrderKey{"USD", "BUY"}][0] {
		t.Errorf("DispatchOrder = %v, want %v", book_in, book_out)
	}
}

func TestPartialSell(t *testing.T) {
	order_in := Order{1396369141, "SELL", 100, "USD", 100, 1000}
	book_in := make(map[OrderKey][]Order)
	book_out := make(map[OrderKey][]Order)
	book_in[OrderKey{"USD", "BUY"}] = []Order{{1396369140, "BUY", 100, "USD", 60, 1001}}
	book_out[OrderKey{"USD", "SELL"}] = []Order{{1396369141, "SELL", 100, "USD", 40, 1000}}
	DispatchOrder(&order_in, &book_in)
	if book_in[OrderKey{"USD", "SELL"}][0] != book_out[OrderKey{"USD", "SELL"}][0] {
		t.Errorf("DispatchOrder = %v, want %v", book_in, book_out)
	}
}
