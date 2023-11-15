package coincheck

type Order struct {
	client *Client
}

// Create a order object with given parameters.In live mode, this issues a transaction.
func (a Order) Create(param string) (string, error) {
	return a.client.Request("POST", "api/exchange/orders", param)
}

// cancel a created order specified by order id. Optional argument amount is to refund partially.
func (a Order) Cancel(id string) (string, error) {
	return a.client.Request("DELETE", "api/exchange/orders/"+id, "")
}

// List charges filtered by params
func (a Order) Opens() (string, error) {
	return a.client.Request("GET", "api/exchange/orders/opens", "")
}

// Get Order Transactions
func (a Order) Transactions() (string, error) {
	return a.client.Request("GET", "api/exchange/orders/transactions", "")
}
