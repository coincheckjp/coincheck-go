package coincheck

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	accessKey    string
	secretKey    string
	account      Account
	bank_account BankAccount
	borrow       Borrow
	deposit      Deposit
	leverage     Leverage
	order        Order
	order_book   OrderBook
	send         Send
	ticker       Ticker
	trade        Trade
	transfer     Transfer
	withdraw     Withdraw
}

func NewClient(accessKey string, secretKey string) *Client {
	c := new(Client)
	c.accessKey = accessKey
	c.secretKey = secretKey
	return c
}

func (c *Client) Request(method string, path string, param string) string {
	if param != "" && method == "GET" {
		path = path + "?" + param
		param = ""
	}
	url := "https://coincheck.jp/" + path
	nonce := strconv.FormatInt(CreateNonce(), 10)
	message := nonce + url + param
	req := &http.Request{}
	if method == "POST" {
		payload := strings.NewReader(param)
		req, _ = http.NewRequest(method, url, payload)
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	signature := ComputeHmac256(message, c.secretKey)
	req.Header.Add("access-key", c.accessKey)
	req.Header.Add("access-nonce", nonce)
	req.Header.Add("access-signature", signature)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	return string(body)
}

// create nonce by milliseconds
func CreateNonce() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// create signature
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) GetAccount() *Account {
	return &Account{
		client: c,
	}
}

func (c *Client) GetBankAccount() *BankAccount {
	return &BankAccount{
		client: c,
	}
}

func (c *Client) GetBorrow() *Borrow {
	return &Borrow{
		client: c,
	}
}

func (c *Client) GetDeposit() *Deposit {
	return &Deposit{
		client: c,
	}
}

func (c *Client) GetLeverage() *Leverage {
	return &Leverage{
		client: c,
	}
}

func (c *Client) GetOrder() *Order {
	return &Order{
		client: c,
	}
}

func (c *Client) GetOrderBook() *OrderBook {
	return &OrderBook{
		client: c,
	}
}

func (c *Client) GetSend() *Send {
	return &Send{
		client: c,
	}
}

func (c *Client) GetTicker() *Ticker {
	return &Ticker{
		client: c,
	}
}

func (c *Client) GetTrade() *Trade {
	return &Trade{
		client: c,
	}
}

func (c *Client) GetTransfer() *Transfer {
	return &Transfer{
		client: c,
	}
}

func (c *Client) GetWithdraw() *Withdraw {
	return &Withdraw{
		client: c,
	}
}
