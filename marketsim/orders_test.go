package marketsim

import "testing"

func TestParseOrder(t *testing.T) {
	in := "2014-04-01T16:19:00Z BUY 100.000 USD 100"
	out := Order{1396369140, "BUY", 100, "USD", 100}
	if x := ParseOrder(in); *x != out {
		t.Errorf("ParseOrder(%v) = %v, want %v", in, x, out)
	}
}

func TestDispatchOrder(t *testing.T) {
    order_in := Order{1396369140, "BUY", 100, "USD", 100}
    book_in := make(map[OrderKey][]Order)
    book_out := make(map[OrderKey][]Order)
    book_out[OrderKey{"USD", "BUY"}] = []Order{{1396369140, "BUY", 100, "USD", 100}}
    DispatchOrder(&order_in, &book_in)
    if book_in[OrderKey{"USD", "BUY"}][0] != book_out[OrderKey{"USD", "BUY"}][0] {
        t.Errorf("DispatchOrder = %v, want %v", book_in, book_out)
    }
}
