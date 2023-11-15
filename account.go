package coincheck

type Account struct {
	client *Client
}

// Make sure a balance.
func (a Account) Balance() (string, error) {
	return a.client.Request("GET", "api/accounts/balance", "")
}

// Make sure a leverage balance.
func (a Account) LeverageBalance() (string, error) {
	return a.client.Request("GET", "api/accounts/leverage_balance", "")
}

// Get account information.
func (a Account) Info() (string, error) {
	return a.client.Request("GET", "api/accounts/balance", "")
}
