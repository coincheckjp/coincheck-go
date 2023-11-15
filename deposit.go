package coincheck

type Deposit struct {
	client *Client
}

// You Get Deposit history
func (a Deposit) All(param string) (string, error) {
	return a.client.Request("GET", "/api/deposit_money", param)
}

// Deposit Bitcoin Faster
func (a Deposit) Fast(id string) (string, error) {
	return a.client.Request("POST", "/api/deposit_money/"+id+"/fast", `{"id":"`+id+`"}`)
}
