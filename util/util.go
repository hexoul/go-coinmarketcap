// Package util supports specific parsing
package util

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/hexoul/go-coinmarketcap/types"
)

// ToInt helper for parsing strings to int
func ToInt(rawInt string) int {
	parsed, _ := strconv.Atoi(strings.Replace(strings.Replace(rawInt, "$", "", -1), ",", "", -1))
	return parsed
}

// ToFloat helper for parsing strings to float
func ToFloat(rawFloat string) float64 {
	parsed, _ := strconv.ParseFloat(strings.Replace(strings.Replace(strings.Replace(rawFloat, "$", "", -1), ",", "", -1), "%", "", -1), 64)
	return parsed
}

// ParseOptions returns parsed param list
func ParseOptions(options *types.Options) (params []string) {
	if options == nil {
		return
	}
	if options.ID != "" {
		params = append(params, fmt.Sprintf("id=%s", options.ID))
	}
	if options.Symbol != "" {
		params = append(params, fmt.Sprintf("symbol=%s", options.Symbol))
	}
	if options.Slug != "" {
		params = append(params, fmt.Sprintf("slug=%s", options.Slug))
	}
	if options.Start != 0 {
		params = append(params, fmt.Sprintf("start=%d", options.Start))
	}
	if options.Limit != 0 {
		params = append(params, fmt.Sprintf("limit=%d", options.Limit))
	}
	if options.Convert != "" {
		params = append(params, fmt.Sprintf("convert=%s", options.Convert))
	}
	if options.Sort != "" {
		params = append(params, fmt.Sprintf("sort=%s", options.Sort))
	}
	if options.SortDir != "" {
		params = append(params, fmt.Sprintf("sort_dir=%s", options.SortDir))
	}
	if options.CryptoType != "" {
		params = append(params, fmt.Sprintf("cryptocurrency_type=%s", options.CryptoType))
	}
	if options.MarketType != "" {
		params = append(params, fmt.Sprintf("market_type=%s", options.MarketType))
	}
	if options.TimePeriod != "" {
		params = append(params, fmt.Sprintf("time_period=%s", options.TimePeriod))
	}
	if options.TimeStart != "" {
		params = append(params, fmt.Sprintf("time_start=%s", options.TimeStart))
	}
	if options.TimeEnd != "" {
		params = append(params, fmt.Sprintf("time_end=%s", options.TimeEnd))
	}
	if options.Interval != "" {
		params = append(params, fmt.Sprintf("interval=%s", options.Interval))
	}
	if options.Count != 0 {
		params = append(params, fmt.Sprintf("interval=%d", options.Count))
	}
	return
}

// DoReq HTTP client
func DoReq(req *http.Request) (body []byte, err error) {
	requestTimeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: requestTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}
	return
}

// InvokeChromedp for scraping AJAX page
func InvokeChromedp(url, qeury string, sec int, buffer *string) (err error) {
	// Create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// NOTE: not appeared error form => chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	dp, err := chromedp.New(ctxt)
	if err != nil {
		return err
	}

	// Run task list
	if err = dp.Run(ctxt, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(time.Duration(sec) * time.Second),
		chromedp.Text(qeury, buffer, chromedp.ByID),
	}); err != nil {
		return
	}

	// Shutdown chrome
	if err = dp.Shutdown(ctxt); err != nil {
		return
	}

	// Wait for chrome to finish
	err = dp.Wait()
	return
}

// ISODate returns ISO date format like "2018-11-23"
func ISODate(t time.Time) string {
	return strings.Split(t.Format(time.RFC3339), "T")[0]
}
