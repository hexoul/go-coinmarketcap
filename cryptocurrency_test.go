// NOTE: before testing, apply your CMC API key to init() of `coinmarketcap_test.go`
package coinmarketcap

import (
	"testing"

	"github.com/hexoul/go-coinmarketcap/types"
)

func TestCryptoInfo(t *testing.T) {
	if ret, err := GetInstance().CryptoInfo(&types.Options{
		Symbol: "BTC",
	}); err != nil {
		t.Fatal(err)
	} else if ret.CryptoInfo["BTC"].Name != "Bitcoin" {
		t.FailNow()
	}
}

func TestCryptoMap(t *testing.T) {
	if ret, err := GetInstance().CryptoMap(&types.Options{
		Symbol: "BTC",
	}); err != nil {
		t.Fatal(err)
	} else if len(ret.CryptoMap) == 0 || ret.CryptoMap[0].Name != "Bitcoin" {
		t.FailNow()
	}
}

func TestCryptoListingsLatest(t *testing.T) {
	if listings, err := GetInstance().CryptoListingsLatest(&types.Options{
		Limit: 1,
	}); err != nil {
		t.Fatal(err)
	} else if len(listings.CryptoMarket) == 0 {
		t.FailNow()
	} else if listings.CryptoMarket[0].Name != "Bitcoin" {
		t.FailNow()
	} else if listings.CryptoMarket[0].Quote["USD"].Price <= 0 {
		t.FailNow()
	}
}

func TestCryptoMarketPairsLatest(t *testing.T) {
	if info, err := GetInstance().CryptoMarketPairsLatest(&types.Options{
		Symbol: "BTC",
	}); err != nil {
		t.Fatal(err)
	} else if info.Name != "Bitcoin" {
		t.FailNow()
	}
}

func TestCryptoOhlcvLatest(t *testing.T) {
	if info, err := GetInstance().CryptoOhlcvLatest(&types.Options{
		Symbol:  "BTC",
		Convert: "USD",
	}); err != nil {
		t.Fatal(err)
	} else if info.Ohlcv["BTC"].Name != "Bitcoin" {
		t.FailNow()
	} else if info.Ohlcv["BTC"].Quote["USD"] == nil {
		t.FailNow()
	}
}

func TestCryptoOhlcvHistorical(t *testing.T) {
	if info, err := GetInstance().CryptoOhlcvHistorical(&types.Options{
		Symbol:    "BTC",
		Convert:   "USD",
		TimeStart: "2018-11-12",
		TimeEnd:   "2018-11-13",
		Interval:  types.IntervalOptions.Hourly,
	}); err != nil {
		t.Fatal(err)
	} else if info.Name != "Bitcoin" {
		t.FailNow()
	} else if info.Ohlcv[0].Quote["USD"] == nil {
		t.FailNow()
	} else if info.Ohlcv[0].Quote["USD"].Volume == 0 {
		t.FailNow()
	}
}

func TestCryptoMarketQuotesLatest(t *testing.T) {
	if quotes, err := GetInstance().CryptoMarketQuotesLatest(&types.Options{
		Symbol: "BTC,ETH",
	}); err != nil {
		t.Fatal(err)
	} else if len(quotes.CryptoMarket) == 0 || len(quotes.CryptoMarket) != 2 {
		t.FailNow()
	} else if quotes.CryptoMarket["BTC"].Name != "Bitcoin" || quotes.CryptoMarket["ETH"].Name != "Ethereum" {
		t.FailNow()
	}
}
