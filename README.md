# coincheck-go

CoinCheck APIs for Go

## Usage

```go
// 残高
client := coincheck.NewClient("ACCESS_KEY", "API_SECRET")
client.GetAccount.Balance()
```

## Installation

```
go get github.com/coincheckjp/coincheck-go
```

installation for CLI

```
go install github.com/coincheckjp/coincheck-go/cmd/coincheck@latest
```
