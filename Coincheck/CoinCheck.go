package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"time"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

type CoinCheck struct {
	accessKey string
	secretKey string
	account Account
}

func (c *CoinCheck) Init(accessKey string, secretKey string) {
	c.accessKey = accessKey
	c.secretKey = secretKey
	c.account = Account{}
}

func Request(method string, path string, param string ) {
	url := "https://coincheck.jp" + path
	nonce := strconv.FormatInt(CreateNonce(), 10)
	message := nonce + url + param
	payload := strings.NewReader(param)
	req, _ := http.NewRequest(method, url, payload)
	signature := ComputeHmac256(message, "DmkP3ChVjOjqieaX-jFnU4XxpSlSgT3M")
	req.Header.Add("access-key", "CTBjkbYihKmT-IHR")
	req.Header.Add("access-nonce", nonce)
	req.Header.Add("access-signature", signature)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

//create nonce by milliseconds
func CreateNonce() int64{
	return time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond));
}

//create signature
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}


