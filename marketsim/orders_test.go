package marketsim

import "testing"

func TestParseOrder(t *testing.T) {
    in := "2014-04-01T16:19:00Z BUY USD 100.000 100"
    out := Order{1396369140, "BUY", "USD", 100, 100}
    if x:= ParseOrder(in); *x != out {
        t.Errorf("ParseOrder(%v) = %v, want %v", in, x, out)
    }
}

