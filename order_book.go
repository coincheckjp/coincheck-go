package coincheck

type OrderBook struct {
	client *Client
}

// 板情報を取得できます。
func (a OrderBook) All() string {
	return a.client.Request("GET", "api/order_books", "")
}
