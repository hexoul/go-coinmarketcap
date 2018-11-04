// Package util supports specific parsing
package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

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
