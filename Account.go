package main

type Account struct {
	client *CoinCheck
}

// Make sure a balance.
func (a Account) balance() string {
	return a.client.Request("GET", "api/accounts/balance", "")
}

// Make sure a leverage balance.
func (a Account) leverage_balance() string {
	return a.client.Request("GET", "api/accounts/leverage_balance", "")
}

// Get account information.
func (a Account) info() string {
	return a.client.Request("GET", "api/accounts/balance", "")
}
