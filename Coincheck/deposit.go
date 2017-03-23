package main

type Deposit struct {
	client *CoinCheck
}


// You Get Deposit history
func (a Deposit) all(param string) string {
	return a.client.Request("GET", "api/deposit_money", param)
}

// Deposit Bitcoin Faster
func (a Deposit) fast(id string) string {
	return a.client.Request("POST", "api/deposit_money/" + id + "/fast", `{"id":"`+id+`"}`)
}
