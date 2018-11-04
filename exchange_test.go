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
	if info["binance"].Name != "Binance" {
		t.FailNow()
	}
}
