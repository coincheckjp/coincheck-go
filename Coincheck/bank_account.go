package main

type BankAccount struct {
	client *CoinCheck
}

//Create a new BankAccount.
func (a BankAccount) create() string {
	return a.client.Request("POST", "api/accounts/balance", "")
}

// Get account information.
func (a BankAccount) all() string {
	return a.client.Request("GET", "api/bank_accounts", "")
}

//Delete a BankAccount.
func (a BankAccount) delete() string {
	return a.client.Request("GET", "api/bank_accounts", "")
}
