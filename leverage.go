package coincheck

type Leverage struct {
	client *Client
}

// Get a leverage positions list.
func (a Leverage) Positions() (string, error) {
	return a.client.Request("GET", "/api/exchange/leverage/positions", "")
}
