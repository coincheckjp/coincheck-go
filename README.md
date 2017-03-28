# coincheck-go

# Install
Go to CoinCheck folder and run this command ```go run *.go```

# API
```go
client := new(CoinCheck).NewClient("ACCESS_KEY", "API_SECRET")

/** Public API */
client.ticker.all()
client.trade.all()
client.order_book.all()

/** Private API */
// 新規注文
// "buy" 指値注文 現物取引 買い
// "sell" 指値注文 現物取引 売り
// "market_buy" 成行注文 現物取引 買い
// "market_sell" 成行注文 現物取引 売り
// "leverage_buy" 指値注文 レバレッジ取引新規 買い
// "leverage_sell" 指値注文 レバレッジ取引新規 売り
// "close_long" 指値注文 レバレッジ取引決済 売り
// "close_short" 指値注文 レバレッジ取引決済 買い
client.order.create(`{"rate":"28500","amount":"0.00508771", "order_type":"buy", "pair":"btc_jpy"}`)

// 未決済の注文一覧
client.order.opens()

// 注文のキャンセル
client.order.cancel("12345")

// 取引履歴
client.order.transactions()

// ポジション一覧
client.leverage.positions()

// 残高
client.account.balance()

// レバレッジアカウントの残高
client.account.leverage_balance()

// アカウント情報
client.account.info()

// ビットコインの送金
client.send.create(`{"address":"1Gp9MCp7FWqNgaUWdiUiRPjGqNVdqug2hY","amount":"0.0002"`)

// ビットコインの送金履歴
client.send.all("currency=BTC")

// ビットコインの受け取り履歴
client.deposit.all("currency=BTC")

// ビットコインの高速入金
client.deposit.fast("12345")

// 銀行口座一覧
client.bank_account.all()

// 銀行口座の登録
client.bank_account.create(`{"bank_name":"MUFG","branch_name":"tokyo", "bank_account_type":"toza", "number":"1234567", "name":"Danny"}`)

// 銀行口座の削除
client.bank_account.delete("25621")

// 出金履歴
client.withdraw.all()

// 出金申請の作成
client.withdraw.create(`{"bank_account_id":"2222","amount":"50000", "currency":"JPY", "is_fast":"false"}`)

// 出金申請のキャンセル
client.withdraw.cancel("12345")

// 借入申請
client.borrow.create(`{"amount":"100","currency":"JPY"}`)

// 借入中一覧
client.borrow.matches()

// 返済
client.borrow.repay("1135")

// レバレッジアカウントへの振替
client.transfer.to_leverage(`{"amount":"100","currency":"JPY"}`)

// レバレッジアカウントからの振替
client.transfer.from_leverage(`{"amount":"100","currency":"JPY"}`)
```
