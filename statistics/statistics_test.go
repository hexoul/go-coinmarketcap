package statistics

import (
	"testing"
	"time"

	"github.com/jasonlvhit/gocron"

	coinmarketcap "github.com/hexoul/go-coinmarketcap"
	"github.com/hexoul/go-coinmarketcap/types"
)

func init() {
	coinmarketcap.GetInstanceWithKey("YOUR_API_KEY")
}

func TestLog(t *testing.T) {
	testLog()
}

func TestCron(t *testing.T) {
	testCron()
	time.Sleep(20 * time.Second)
}

func TestGatherCryptoQuote(t *testing.T) {
	GatherCryptoQuote(&types.Options{
		Symbol: "BTC",
	}, gocron.Every(10).Seconds())
	gocron.Start()
	time.Sleep(30 * time.Second)
}

func TestGatherTokenMetric(t *testing.T) {
	GatherTokenMetric("BNB", "0xB8c77482e45F1F44dE1745F52C74426C631bDD52", gocron.Every(10).Seconds())
	gocron.Start()
	time.Sleep(30 * time.Second)
}
