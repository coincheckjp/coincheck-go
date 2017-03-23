package main

type Order struct {
	client *CoinCheck
}

// Create a order object with given parameters.In live mode, this issues a transaction.
func (a Order) create(param string) string {
	return a.client.Request("POST", "api/exchange/orders", param)
}

// cancel a created order specified by order id. Optional argument amount is to refund partially.
func (a Order) cancel(id string) string {
	return a.client.Request("DELETE", "api/exchange/orders/" + id, "")
}

// List charges filtered by params
func (a Order) opens() string {
	return a.client.Request("GET", "api/exchange/orders/opens", "")
}

// Get Order Transactions
func (a Order) transactions() string {
	return a.client.Request("GET", "api/exchange/orders/transactions", "")
}