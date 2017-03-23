package main

type Trade struct {
	client *CoinCheck
}

// 最新の取引履歴を取得できます。
func (a Trade) all() string {
	return a.client.Request("GET", "api/trades", "")
}
