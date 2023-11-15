package main

import (
	"github.com/coincheckjp/coincheck-go"
)

func main() {
	client := coincheck.NewClient("CTBjkbYihKmT-IHR", "DmkP3ChVjOjqieaX-jFnU4XxpSlSgT3M")
	/** Public API */
	client.GetTicker().All()
	client.GetTrade().All()
	client.GetOrderBook().All()

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
	client.GetOrder().Create(`{"rate":"28500","amount":"0.00508771", "order_type":"buy", "pair":"btc_jpy"}`)
	// 未決済の注文一覧
	client.GetOrder().Opens()
	// 注文のキャンセル
	client.GetOrder().Cancel("12345")
	// 取引履歴
	client.GetOrder().Transactions()
	// ポジション一覧
	client.GetLeverage().Positions()
	// 残高
	client.GetAccount().Balance()
	// レバレッジアカウントの残高
	client.GetAccount().LeverageBalance()
	// アカウント情報
	client.GetAccount().Info()
	// ビットコインの送金
	client.GetSend().Create(`{"address":"1Gp9MCp7FWqNgaUWdiUiRPjGqNVdqug2hY","amount":"0.0002"`)
	// ビットコインの送金履歴
	client.GetSend().All("currency=BTC")
	// ビットコインの受け取り履歴
	client.GetDeposit().All("currency=BTC")
	// ビットコインの高速入金
	client.GetDeposit().Fast("12345")
	// 銀行口座一覧
	client.GetBankAccount().All()
	// 銀行口座の登録
	client.GetBankAccount().Create(`{"bank_name":"MUFG","branch_name":"tokyo", "bank_account_type":"toza", "number":"1234567", "name":"Danny"}`)
	// 銀行口座の削除
	client.GetBankAccount().Delete("25621")
	// 出金履歴
	client.GetWithdraw().All()
	// 出金申請の作成
	client.GetWithdraw().Create(`{"bank_account_id":"2222","amount":"50000", "currency":"JPY", "is_fast":"false"}`)
	// 出金申請のキャンセル
	client.GetWithdraw().Cancel("12345")
	// 借入申請
	client.GetBorrow().Create(`{"amount":"100","currency":"JPY"}`)
	// 借入中一覧
	client.GetBorrow().Matches()
	// 返済
	client.GetBorrow().Repay("1135")
	// レバレッジアカウントへの振替
	client.GetTransfer().ToLeverage(`{"amount":"100","currency":"JPY"}`)
	// レバレッジアカウントからの振替
	client.GetTransfer().FromLeverage(`{"amount":"100","currency":"JPY"}`)
}
