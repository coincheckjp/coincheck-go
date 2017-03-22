package main

type Ticker struct {
	client *CoinCheck
}

// 各種最新情報を簡易に取得することができます。
func (a Ticker) all(param string) string {
	return a.client.Request("GET", "api/ticker", "")
}
