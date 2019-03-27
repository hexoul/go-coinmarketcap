# go-coinmarketcap

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hexoul/go-coinmarketcap/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hexoul/go-coinmarketcap)](https://goreportcard.com/report/github.com/hexoul/go-coinmarketcap)
[![GoDoc](https://godoc.org/github.com/hexoul/go-coinmarketcap?status.svg)](https://godoc.org/github.com/hexoul/go-coinmarketcap)

> Coinmarketcap (CMC) Pro API Client written in Golang

## Install

`go get -u github.com/hexoul/go-coinmarketcap`

## Getting started

- As library, start from `coinmarketcap.GetInstanceWithKey('YOUR_API_KEY')`
- As program, start from `coinmarketcap.GetInstance()` after executing `go run -cmcApikey=[YOUR_API_KEY]`

```go
package main

import (
    "fmt"

    cmc "github.com/hexoul/go-coinmarketcap"
    "github.com/hexoul/go-coinmarketcap/types"
)

func init() {
    cmc.GetInstanceWithKey("YOUR_API_KEY")
}

func main() {
    if crypto, err := cmc.GetInstance().CryptoInfo(&types.Options{
        Symbol: "BTC",
    }); err == nil {
        fmt.Println(crypto.CryptoInfo["BTC"].Name)
    }

    client := cmc.GetInstance()
    if exchange, err := client.ExchangeInfo(&types.Options{
        Slug: "binance",
    }); err == nil {
        fmt.Println(exchange.ExchangeInfo["binance"].Name)
    }
}
```

## Features

| Type           | Endpoint                               | Done |
|----------------|----------------------------------------|-------------|
| Cryptocurrency | /v1/cryptocurrency/info                | ✔ |
| Cryptocurrency | /v1/cryptocurrency/map                 | ✔ |
| Cryptocurrency | /v1/cryptocurrency/listings/latest     | ✔ |
| Cryptocurrency | /v1/cryptocurrency/listings/historical | - |
| Cryptocurrency | /v1/cryptocurrency/market-pairs/latest | ✔ |
| Cryptocurrency | /v1/cryptocurrency/ohlcv/latest        | ✔ |
| Cryptocurrency | /v1/cryptocurrency/ohlcv/historical    | ✔ |
| Cryptocurrency | /v1/cryptocurrency/quotes/latest       | ✔ |
| Cryptocurrency | /v1/cryptocurrency/quotes/historical   | - |
| Exchange       | /v1/exchange/info                      | ✔ |
| Exchange       | /v1/exchange/map                       | ✔ |
| Exchange       | /v1/exchange/listings/latest           | ✔ |
| Exchange       | /v1/exchange/listings/historical       | - |
| Exchange       | /v1/exchange/market-pairs/latest       | ✔ |
| Exchange       | /v1/exchange/quotes/latest             | ✔ |
| Exchange       | /v1/exchange/quotes/historical         | - |
| Global Metrics | /v1/global-metrics/quotes/latest       | - |
| Global Metrics | /v1/global-metrics/quotes/historical   | - |
| Tools          | /v1/tools/price-conversion             | - |

## Reference

[Coinmarketcap (CMC) Pro](https://pro.coinmarketcap.com/api/v1)
