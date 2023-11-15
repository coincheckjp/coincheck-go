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
						result, err := client.GetTicker().All()
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetAccount().Balance()
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "leveragebalance",
					Usage: "show leverage balances",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// レバレッジアカウントの残高
						result, err := client.GetAccount().LeverageBalance()
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "info",
					Usage: "show info",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// アカウント情報
						result, err := client.GetAccount().Info()
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetBankAccount().All()
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "create",
					Usage: "create bank account",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 銀行口座の登録
						result, err := client.GetBankAccount().Create(`{"bank_name":"MUFG","branch_name":"tokyo", "bank_account_type":"toza", "number":"1234567", "name":"Danny"}`)
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "delete bank account",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 銀行口座の削除
						result, err := client.GetBankAccount().Delete("25621")
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetSend().All("currency=BTC")
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "create",
					Usage: "create send",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// ビットコインの送金
						result, err := client.GetSend().Create(`{"address":"1Gp9MCp7FWqNgaUWdiUiRPjGqNVdqug2hY","amount":"0.0002"`)
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetDeposit().All("currency=BTC")
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "fast",
					Usage: "payment fast",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// ビットコインの高速入金
						result, err := client.GetDeposit().Fast("12345")
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetTrade().All()
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetOrder().Create(`{"rate":"28500","amount":"0.00508771", "order_type":"buy", "pair":"btc_jpy"}`)
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "cancel",
					Usage: "cancel order",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 注文のキャンセル
						result, err := client.GetOrder().Cancel("12345")
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "opens",
					Usage: "show all orders opened",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 未決済の注文一覧
						result, err := client.GetOrder().Opens()
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "transactions",
					Usage: "show all transactions",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 取引履歴
						result, err := client.GetOrder().Transactions()
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetOrderBook().All()
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetLeverage().Positions()
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetWithdraw().All()
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "create",
					Usage: "create withdraw",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 出金申請の作成
						result, err := client.GetWithdraw().Create(`{"bank_account_id":"2222","amount":"50000", "currency":"JPY", "is_fast":"false"}`)
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "cancel",
					Usage: "cancel withdraw",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 出金申請のキャンセル
						result, err := client.GetWithdraw().Cancel("12345")
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetBorrow().Create(`{"amount":"100","currency":"JPY"}`)
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "matches",
					Usage: "show all matches",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 借入中一覧
						result, err := client.GetBorrow().Matches()
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "repay",
					Usage: "repay borrow",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 返済
						result, err := client.GetBorrow().Repay("1135")
						if err != nil {
							return err
						}
						fmt.Println(result)
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
						result, err := client.GetTransfer().ToLeverage(`{"amount":"100","currency":"JPY"}`)
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "from-leverage",
					Usage: "transfer from leverage",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// レバレッジアカウントからの振替
						result, err := client.GetTransfer().FromLeverage(`{"amount":"100","currency":"JPY"}`)
						if err != nil {
							return err
						}
						fmt.Println(result)
						return nil
					},
				},
				{
					Name:  "repay",
					Usage: "repay borrow",
					Action: func(cCtx *cli.Context) error {
						client := cCtx.App.Metadata["client"].(*coincheck.Client)
						// 返済
						result, err := client.GetBorrow().Repay("1135")
						if err != nil {
							return err
						}
						fmt.Println(result)
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
