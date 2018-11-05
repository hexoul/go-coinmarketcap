package coinmarketcap

import (
	"testing"

	"github.com/hexoul/go-coinmarketcap/types"
)

func init() {
	GetInstanceWithKey("YOUR_API_KEY")
}

func TestExchangeInfo(t *testing.T) {
	info, err := GetInstance().ExchangeInfo(&types.Options{
		Slug: "binance",
	})
	if err != nil {
		t.Fatal(err)
	}
	if info.ExchangeInfo["binance"].Name != "Binance" {
		t.FailNow()
	}
}

func TestExchangeMap(t *testing.T) {
	info, err := GetInstance().ExchangeMap(&types.Options{
		Slug: "binance",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(info.ExchangeMap) == 0 {
		t.FailNow()
	}
	if info.ExchangeMap[0].Name != "Binance" {
		t.FailNow()
	}
}

func TestExchangeListingsLatest(t *testing.T) {
	info, err := GetInstance().ExchangeListingsLatest(&types.Options{
		Limit: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(info.MarketQuote) == 0 || len(info.MarketQuote) == 2 {
		t.FailNow()
	}
	if info.MarketQuote[0].Name != "Binance" {
		t.FailNow()
	}
}

func TestExchangeMarketPairsLatest(t *testing.T) {
	info, err := GetInstance().ExchangeMarketPairsLatest(&types.Options{
		Slug: "binance",
	})
	if err != nil {
		t.Fatal(err)
	}
	if info.Name != "Binance" {
		t.FailNow()
	}
}

func TestExchangeMarketQuotesLatest(t *testing.T) {
	info, err := GetInstance().ExchangeMarketQuotesLatest(&types.Options{
		Slug: "binance",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(info.MarketQuote) == 0 {
		t.FailNow()
	}
	if info.MarketQuote["binance"].Name != "Binance" {
		t.FailNow()
	}
}
