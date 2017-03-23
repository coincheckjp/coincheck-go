package main

type Withdraw struct {
	client *CoinCheck
}

// Transfer Balance to Leverage.
func (a Withdraw) create(param string) string {
	return a.client.Request("POST", "api/withdraws", param)
}

// Transfer Balance from Leverage.
func (a Withdraw) all() string {
	return a.client.Request("GET", "api/withdraws", "")
}

// Transfer Balance from Leverage.
func (a Withdraw) cancel(id string) string {
	return a.client.Request("DELETE", "api/withdraws/" + id, "")
}