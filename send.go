package coincheck

type Send struct {
	client *Client
}

// Sending Bitcoin to specified Bitcoin addres.
func (a Send) Create(param string) string {
	return a.client.Request("POST", "api/send_money", param)
}

// You Get Send history
func (a Send) All(param string) string {
	return a.client.Request("GET", "api/send_money", param)
}
