package main

import (
	"fmt"
	"github.com/coincheckjp/coincheck-go"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var app = &cli.App{
	Name:  "coincheck",
	Usage: "coincheck [command]",
	Action: func(*cli.Context) error {
		return nil
	},
	Commands: []*cli.Command{
		{
			Name:  "ticker",
			Usage: "operations for ticker",
			Subcommands: []*cli.Command{
				{
					Name:  "all",
					Usage: "show all tickers",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						fmt.Println(client.GetTicker().All())
						return nil
					},
				},
			},
		},
		{
			Name:  "account",
			Usage: "operations for account",
			Subcommands: []*cli.Command{
				{
					Name:  "balance",
					Usage: "show balances",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 残高
						client.GetAccount().Balance()
						return nil
					},
				},
				{
					Name:  "leveragebalance",
					Usage: "show leverage balances",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// レバレッジアカウントの残高
						client.GetAccount().LeverageBalance()
						return nil
					},
				},
				{
					Name:  "info",
					Usage: "show info",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// アカウント情報
						client.GetAccount().Info()
						return nil
					},
				},
			},
		},
		{
			Name:  "bank-account",
			Usage: "operations for bank account",
			Subcommands: []*cli.Command{
				{
					Name:  "all",
					Usage: "show all bank accounts",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 銀行口座一覧
						client.GetBankAccount().All()
						return nil
					},
				},
				{
					Name:  "create",
					Usage: "create bank account",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 銀行口座の登録
						client.GetBankAccount().Create(`{"bank_name":"MUFG","branch_name":"tokyo", "bank_account_type":"toza", "number":"1234567", "name":"Danny"}`)
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "delete bank account",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 銀行口座の削除
						client.GetBankAccount().Delete("25621")
						return nil
					},
				},
			},
		},
		{
			Name:  "send",
			Usage: "operations for send",
			Subcommands: []*cli.Command{
				{
					Name:  "all",
					Usage: "show all sends",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// ビットコインの送金履歴
						client.GetSend().All("currency=BTC")
						return nil
					},
				},
				{
					Name:  "create",
					Usage: "create send",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// ビットコインの送金
						client.GetSend().Create(`{"address":"1Gp9MCp7FWqNgaUWdiUiRPjGqNVdqug2hY","amount":"0.0002"`)
						return nil
					},
				},
			},
		},
		{
			Name:  "deposit",
			Usage: "operations for deposit",
			Subcommands: []*cli.Command{
				{
					Name:  "all",
					Usage: "show all deposits",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// ビットコインの受け取り履歴
						client.GetDeposit().All("currency=BTC")
						return nil
					},
				},
				{
					Name:  "fast",
					Usage: "payment fast",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// ビットコインの高速入金
						client.GetDeposit().Fast("12345")
						return nil
					},
				},
			},
		},
		{
			Name:  "trade",
			Usage: "operations for trade",
			Subcommands: []*cli.Command{
				{
					Name:  "all",
					Usage: "show all trades",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						fmt.Println(client.GetTrade().All())
						return nil
					},
				},
			},
		},
		{
			Name:  "order",
			Usage: "operations for order",
			Subcommands: []*cli.Command{
				{
					Name:  "create",
					Usage: "create order",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 注文の作詞絵
						client.GetOrder().Create(`{"rate":"28500","amount":"0.00508771", "order_type":"buy", "pair":"btc_jpy"}`)
						return nil
					},
				},
				{
					Name:  "cancel",
					Usage: "cancel order",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 注文のキャンセル
						client.GetOrder().Cancel("12345")
						return nil
					},
				},
				{
					Name:  "opens",
					Usage: "show all orders opened",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 未決済の注文一覧
						client.GetOrder().Opens()
						return nil
					},
				},
				{
					Name:  "transactions",
					Usage: "show all transactions",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 取引履歴
						client.GetOrder().Transactions()
						return nil
					},
				},
			},
		},
		{
			Name:  "orderbook",
			Usage: "operations for orderbook",
			Subcommands: []*cli.Command{
				{
					Name:  "all",
					Usage: "show all orderbooks",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						fmt.Println(client.GetOrderBook().All())
						return nil
					},
				},
			},
		},
		{
			Name:  "leverage",
			Usage: "operations for leverage",
			Subcommands: []*cli.Command{
				{
					Name:  "position",
					Usage: "show all positions",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// ポジション一覧
						client.GetLeverage().Positions()
						return nil
					},
				},
			},
		},
		{
			Name:  "withdraw",
			Usage: "operations for withdraw",
			Subcommands: []*cli.Command{
				{
					Name:  "all",
					Usage: "show all withdraws",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 出金履歴
						client.GetWithdraw().All()
						return nil
					},
				},
				{
					Name:  "create",
					Usage: "create withdraw",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 出金申請の作成
						client.GetWithdraw().Create(`{"bank_account_id":"2222","amount":"50000", "currency":"JPY", "is_fast":"false"}`)
						return nil
					},
				},
				{
					Name:  "cancel",
					Usage: "cancel withdraw",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 出金申請のキャンセル
						client.GetWithdraw().Cancel("12345")
						return nil
					},
				},
			},
		},
		{
			Name:  "borrow",
			Usage: "operations for borrow",
			Subcommands: []*cli.Command{
				{
					Name:  "create",
					Usage: "create borrow",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 借入申請
						client.GetBorrow().Create(`{"amount":"100","currency":"JPY"}`)
						return nil
					},
				},
				{
					Name:  "matches",
					Usage: "show all matches",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 借入中一覧
						client.GetBorrow().Matches()
						return nil
					},
				},
				{
					Name:  "repay",
					Usage: "repay borrow",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 返済
						client.GetBorrow().Repay("1135")
						return nil
					},
				},
			},
		},
		{
			Name:  "transfer",
			Usage: "operations for transfer",
			Subcommands: []*cli.Command{
				{
					Name:  "to-leverage",
					Usage: "transfer to leverage",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// レバレッジアカウントへの振替
						client.GetTransfer().ToLeverage(`{"amount":"100","currency":"JPY"}`)
						return nil
					},
				},
				{
					Name:  "from-leverage",
					Usage: "transfer from leverage",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// レバレッジアカウントからの振替
						client.GetTransfer().FromLeverage(`{"amount":"100","currency":"JPY"}`)
						return nil
					},
				},
				{
					Name:  "repay",
					Usage: "repay borrow",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 返済
						client.GetBorrow().Repay("1135")
						return nil
					},
				},
			},
		},
	},
	Metadata: map[string]any{},
	Before: func(cCtx *cli.Context) error {
		client := coincheck.NewClient("CTBjkbYihKmT-IHR", "DmkP3ChVjOjqieaX-jFnU4XxpSlSgT3M")
		cCtx.App.Metadata["client"] = client
		return nil
	},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

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
}
