// NOTE: before testing, apply your CMC API key to init() of `coinmarketcap_test.go`
package coinmarketcap

import (
	"testing"

	"github.com/hexoul/go-coinmarketcap/types"
)

func TestExchangeInfo(t *testing.T) {
	if info, err := GetInstance().ExchangeInfo(&types.Options{
		Slug: "binance",
	}); err != nil {
		t.Fatal(err)
	} else if info.ExchangeInfo["binance"].Name != "Binance" {
		t.FailNow()
	}
}

func TestExchangeMap(t *testing.T) {
	if info, err := GetInstance().ExchangeMap(&types.Options{
		Slug: "binance",
	}); err != nil {
		t.Fatal(err)
	} else if len(info.ExchangeMap) == 0 {
		t.FailNow()
	} else if info.ExchangeMap[0].Name != "Binance" {
		t.FailNow()
	}
}

func TestExchangeListingsLatest(t *testing.T) {
	if info, err := GetInstance().ExchangeListingsLatest(&types.Options{
		Limit: 2,
	}); err != nil {
		t.Fatal(err)
	} else if len(info.MarketQuote) < 2 {
		t.FailNow()
	} else if info.MarketQuote[0].Name == "" {
		t.FailNow()
	}
}

func TestExchangeMarketPairsLatest(t *testing.T) {
	if info, err := GetInstance().ExchangeMarketPairsLatest(&types.Options{
		Slug:    "binance",
		Limit:   200,
		Convert: "USD",
	}); err != nil {
		t.Fatal(err)
	} else if info.Name != "Binance" {
		t.FailNow()
	} else {
		for _, pair := range info.MarketPair {
			t.Log(pair.MarketPair, pair.Quote["USD"].Price)
		}
	}
}

func TestExchangeMarketQuotesLatest(t *testing.T) {
	if info, err := GetInstance().ExchangeMarketQuotesLatest(&types.Options{
		Slug: "binance",
	}); err != nil {
		t.Fatal(err)
	} else if len(info.MarketQuote) == 0 {
		t.FailNow()
	} else if info.MarketQuote["binance"].Name != "Binance" {
		t.FailNow()
	}
}
