package coinmarketcap

import (
	"encoding/json"
	"testing"

	"github.com/hexoul/go-coinmarketcap/types"
)

func init() {
	GetInstanceWithKey("YOUR_API_KEY")
}

func TestInfo(t *testing.T) {
	info, err := GetInstance().Info(&types.Options{
		Symbol: "BTC",
	})
	if err != nil {
		t.Error(err)
	}
	if info["BTC"].Name != "Bitcoin" {
		t.FailNow()
	}

	if ret, err := json.Marshal(info); err == nil {
		t.Log(string(ret))
	}
}

func TestListingsLatest(t *testing.T) {
	listings, err := GetInstance().ListingsLatest(&types.Options{
		Limit: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if len(listings) == 0 {
		t.FailNow()
	}
	if listings[0].Name != "Bitcoin" {
		t.FailNow()
	}
	if listings[0].Quote["USD"].Price <= 0 {
		t.FailNow()
	}

	if ret, err := json.Marshal(listings); err == nil {
		t.Log(string(ret))
	}
}
