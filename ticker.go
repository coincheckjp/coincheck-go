package coincheck

type Ticker struct {
	client *Client
}

// 各種最新情報を簡易に取得することができます。
func (a Ticker) all() string {
	return a.client.Request("GET", "api/ticker", "")
}
