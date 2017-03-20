package main

func main() {
	client := new(CoinCheck).NewClient("CTBjkbYihKmT-IHR", "DmkP3ChVjOjqieaX-jFnU4XxpSlSgT3M")
	client.account.balance()
	client.account.leverage_balance()
	client.account.info()
	client.bank_account.all()
	client.bank_account.create()
	client.bank_account.delete()
}