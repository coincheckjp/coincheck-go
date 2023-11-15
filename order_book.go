package main

type OrderBook struct {
	client *CoinCheck
}

// 板情報を取得できます。
func (a OrderBook) all() string {
	return a.client.Request("GET", "api/order_books", "")
}
