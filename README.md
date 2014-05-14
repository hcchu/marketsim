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

  ```map[{OrderKey}][]Order```


