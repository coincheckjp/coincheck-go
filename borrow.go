package coincheck

type Borrow struct {
	client *Client
}

// Create a new Borrow.
func (a Borrow) Create(param string) (string, error) {
	return a.client.Request("POST", "/api/lending/borrows", param)
}

// Get a borrowing list.
func (a Borrow) Matches() (string, error) {
	return a.client.Request("GET", "/api/lending/borrows/matches", "")
}

// Based on this id, you can repay.
func (a Borrow) Repay(id string) (string, error) {
	return a.client.Request("POST", "/api/lending/borrows/"+id+"/repay", `{"id":"`+id+`"}`)
}
