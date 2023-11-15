package coincheck

type Trade struct {
	client *Client
}

// 最新の取引履歴を取得できます。
func (a Trade) All() string {
	return a.client.Request("GET", "api/trades", "")
}
