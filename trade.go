package coincheck

type Trade struct {
	client *Client
}

// 最新の取引履歴を取得できます。
func (a Trade) All() (string, error) {
	return a.client.Request("GET", "api/trades", "")
}
