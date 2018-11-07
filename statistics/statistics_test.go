package statistics

import (
	"testing"
	"time"

	kucoin "github.com/eeonevision/kucoin-go"
	"github.com/jasonlvhit/gocron"

	coinmarketcap "github.com/hexoul/go-coinmarketcap"
	"github.com/hexoul/go-coinmarketcap/types"
)

func init() {
	coinmarketcap.GetInstanceWithKey("YOUR_API_KEY")
}

func TestGatherCryptoQuote(t *testing.T) {
	GatherCryptoQuote(&types.Options{
		Symbol: "BTC",
	}, gocron.Every(10).Seconds())
	gocron.Start()
	time.Sleep(20 * time.Second)
}

func TestGatherTokenMetric(t *testing.T) {
	GatherTokenMetric("BNB", "0xB8c77482e45F1F44dE1745F52C74426C631bDD52", gocron.Every(1).Second())
	gocron.Start()
	time.Sleep(5 * time.Second)
}

func TestGatherKucoinBalance(t *testing.T) {
	k := kucoin.New("API_KEY", "API_SECRET")
	GatherKucoinBalance(k, "BTC", gocron.Every(10).Seconds())
	gocron.Start()
	time.Sleep(20 * time.Second)
}

func TestKucoinListMergedDealtOrders(t *testing.T) {
	k := kucoin.New("API_KEY", "API_SECRET")
	if ret, err := k.ListMergedDealtOrders("ETH-BTC", "BUY", 20, 1, 0, 0); err == nil {
		t.Logf("%v\n", ret)
	}
}
