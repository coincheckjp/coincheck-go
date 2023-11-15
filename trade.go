package coincheck

type Trade struct {
	client *Client
}

// 最新の取引履歴を取得できます。
func (a Trade) all() string {
	return a.client.Request("GET", "api/trades", "")
}
