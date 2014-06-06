## marketsim

An exercise in implementing a simple market/exchange engine in Go.

It currently only supports limit orders, although I may attempt to implement
other order types. Orders come in on STDIN, are parsed into an Order type, and 
then passed onto the matching engine. Orders in the order book are sorted by
price-time priority: more aggressive prices are given priority. In the event of
multiple orders with the same price, earlier orders have priority.

#### Details

* Input timestamps are in RFC3339 format, but are converted to epoch time in
  the order book. A real exchange would be much more precise.

* There's not much in the way of input validation (or error checking) yet.. 

* The order book is a `map` of a struct key that looks like `{"USD", "BUY"}` to
  a slice of Order types.

  ```
  map[OrderKey][]Order
  ```

#### Running

```
$ cat orders.txt
2013-12-01T16:19:00Z    BUY     100.000 BTC     100
2013-12-01T16:19:01Z    BUY     100.000 USD     100
2013-12-01T16:19:10Z    BUY     101.000 BTC     90 
2013-12-01T16:19:11Z    SELL     105.000 BTC     90
2013-12-01T16:19:12Z    SELL     103.000 BTC     90
2013-12-01T16:19:13Z    SELL     104.000 BTC     90
2013-12-01T16:19:13Z    SELL     99.000 USD     90
2013-12-01T16:19:30Z    BUY     110.000 BTC     104 
2013-12-01T16:19:30Z    SELL     110.000 BTC     104
2013-12-01T16:19:35Z    SELL     100.000 BTC     104
2013-12-01T16:19:13Z    BUY     99.000 USD     90
2013-12-01T16:19:36Z    SELL     102.000 BTC     104
2013-12-01T16:19:39Z    BUY     103.000 BTC     104
2013-12-01T16:19:40Z    SELL     100.000 BTC     104

$ go run main.go < orders.txt
SELL ORDER 1006 EXECUTED for 90 USD at 99 matching 1001
BUY ORDER 1007 EXECUTED for 90 BTC at 110 matching 1004
BUY ORDER 1007 EXECUTED for 14 BTC at 110 matching 1005
SELL ORDER 1009 EXECUTED for 90 BTC at 100 matching 1002
SELL ORDER 1009 EXECUTED for 14 BTC at 100 matching 1000
BUY ORDER 1012 EXECUTED for 104 BTC at 103 matching 1011
SELL ORDER 1013 EXECUTED for 86 BTC at 100 matching 1000
EOF
```

