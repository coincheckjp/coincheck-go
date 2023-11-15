package coincheck

type Leverage struct {
	client *Client
}

// Get a leverage positions list.
func (a Leverage) positions() string {
	return a.client.Request("GET", "api/exchange/leverage/positions", "")
}
