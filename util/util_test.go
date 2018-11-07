package util

import (
	"context"
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

func TestDp(t *testing.T) {
	url := "https://etherscan.io/token/0xB8c77482e45F1F44dE1745F52C74426C631bDD52"

	// Initialize dp
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()
	dp, _ := chromedp.New(ctxt)

	var site, res string
	dp.Run(ctxt, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Text(&site, &res),
		chromedp.Sleep(10 * time.Second),
	})
	dp.Shutdown(ctxt)
	dp.Wait()
	fmt.Println(site)
	fmt.Println(res)
}
