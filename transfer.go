package coincheck

type Transfer struct {
	client *Client
}

// Transfer Balance to Leverage.
func (a Transfer) ToLeverage(param string) string {
	return a.client.Request("POST", "api/exchange/transfers/to_leverage", param)
}

// Transfer Balance from Leverage.
func (a Transfer) FromLeverage(param string) string {
	return a.client.Request("POST", "api/exchange/transfers/from_leverage", param)
}
