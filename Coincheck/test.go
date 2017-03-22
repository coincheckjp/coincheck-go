package main

func main() {
	client := new(CoinCheck).NewClient("CTBjkbYihKmT-IHR", "DmkP3ChVjOjqieaX-jFnU4XxpSlSgT3M")
	client.account.balance()
	client.account.leverage_balance()
	client.account.info()
	client.bank_account.all()
	client.bank_account.create(`{"bank_name":"MUFG","branch_name":"tokyo", "bank_account_type":"toza", "number":"1234567", "name":"Danny"}`)
	client.bank_account.delete("25621")
}