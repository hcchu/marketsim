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
    sell_in := []Order{}
    buy_in := make([]Order,0)
    buy_out := []Order{{1396369140, "BUY", 100, "USD", 100}}
    DispatchOrder(&order_in, &buy_in, &sell_in)
    if buy_in[0] != buy_out[0] {
        t.Errorf("DispatchOrder = %v, want %v", buy_in, buy_out)
    }
}
