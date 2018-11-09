package util

import (
	"context"
	"log"
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

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	dp, _ := chromedp.New(ctxt)

	// run task list
	var res string
	dp.Run(ctxt, chromedp.Tasks{
		chromedp.Navigate(`https://etherscan.io/token/0xB8c77482e45F1F44dE1745F52C74426C631bDD52`),
		chromedp.Sleep(10 * time.Second),
		chromedp.Text(`#ContentPlaceHolder1_divSummary td span#totaltxns`, &res, chromedp.ByID),
	})

	// shutdown chrome
	dp.Shutdown(ctxt)

	// wait for chrome to finish
	dp.Wait()

	log.Printf("transfers: %s", res)
}
