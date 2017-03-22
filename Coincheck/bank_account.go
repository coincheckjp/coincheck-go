package main

type BankAccount struct {
	client *CoinCheck
}

// Create a new BankAccount.
func (a BankAccount) create(param string) string {
	return a.client.Request("POST", "api/bank_accounts", param)
}

// Get account information.
func (a BankAccount) all() string {
	return a.client.Request("GET", "api/bank_accounts", "")
}

// Delete a BankAccount.
func (a BankAccount) delete(id string) string {
	return a.client.Request("DELETE", "api/bank_accounts/" + id, "")
}
