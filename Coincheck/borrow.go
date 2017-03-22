package main

type Borrow struct {
	client *CoinCheck
}

// Create a new Borrow.
func (a Borrow) create(param string) string {
	return a.client.Request("POST", "api/lending/borrows", param)
}

// Get a borrowing list.
func (a Borrow) matches() string {
	return a.client.Request("GET", "api/lending/borrows/matches", "")
}

// Based on this id, you can repay.
func (a Borrow) repay(id string) string {
	return a.client.Request("POST", "api/lending/borrows/" + id + "/repay", `{"id":"`+id+`"}`)
}
