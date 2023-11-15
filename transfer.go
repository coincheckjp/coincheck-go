package main

type Transfer struct {
	client *CoinCheck
}

// Transfer Balance to Leverage.
func (a Transfer) to_leverage(param string) string {
	return a.client.Request("POST", "api/exchange/transfers/to_leverage", param)
}

// Transfer Balance from Leverage.
func (a Transfer) from_leverage(param string) string {
	return a.client.Request("POST", "api/exchange/transfers/from_leverage", param)
}
