package util

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/chromedp/chromedp"

	"github.com/hexoul/go-coinmarketcap/types"
)

func TestParseOptions(t *testing.T) {
	if len(ParseOptions(nil)) != 0 {
		t.FailNow()
	}
	if len(ParseOptions(&types.Options{
		Symbol: "BTC",
	})) != 1 {
		t.FailNow()
	}
	if len(ParseOptions(&types.Options{
		Limit: 1,
		Sort:  types.SortOptions.Name,
	})) != 2 {
		t.FailNow()
	}
}

func TestParse(t *testing.T) {
	s := &types.Options{
		Symbol:     "BTC",
		Slug:       "binance",
		Limit:      100,
		CryptoType: "tokens",
	}
	b, err := json.Marshal(s)
	if err != nil {
		t.FailNow()
	}
	p := new(map[string]interface{})
	if err := json.Unmarshal(b, &p); err != nil {
		t.FailNow()
	}
	for k, v := range *p {
		t.Logf(fmt.Sprintf("%s=%v", k, v))
	}
}

func TestDp(t *testing.T) {
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	dp, _ := chromedp.New(ctxt)

	var res string
	dp.Run(ctxt, chromedp.Tasks{
		chromedp.Navigate("https://etherscan.io/token/0xB8c77482e45F1F44dE1745F52C74426C631bDD52"),
		chromedp.Sleep(10 * time.Second),
		chromedp.Text("#ContentPlaceHolder1_divSummary td span#totaltxns", &res, chromedp.ByID),
	})

	dp.Shutdown(ctxt)
	dp.Wait()

	if res == "" {
		t.FailNow()
	}
}
