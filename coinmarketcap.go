// Package coinmarketcap is an API Client for CMC Pro
package coinmarketcap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/hexoul/go-coinmarketcap/types"
	"github.com/hexoul/go-coinmarketcap/util"
)

// Interface for APIs
type Interface interface {
	Info(options *types.InfoOptions) (map[string]*types.Info, error)
	ListingsLatest(options *types.ListingsLatestOptions) ([]*types.Listing, error)
}

// Client the CoinMarketCap client
type Client struct {
	proAPIKey string
}

var (
	instance *Client
	once     sync.Once
	apiKey   string

	// ErrCouldNotCast could not cast error
	ErrCouldNotCast = errors.New("could not cast")
)

const (
	baseURL               = "https://pro-api.coinmarketcap.com/v1"
	coinGraphURL          = "https://graphs2.coinmarketcap.com/currencies"
	globalMarketGraphURL  = "https://graphs2.coinmarketcap.com/global/marketcap-total"
	altcoinMarketGraphURL = "https://graphs2.coinmarketcap.com/global/marketcap-altcoin"
)

func init() {
	for _, val := range os.Args {
		arg := strings.Split(val, "=")
		if len(arg) < 2 {
			continue
		} else if arg[0] == "-apikey" {
			apiKey = arg[1]
		}
	}
}

// GetInstance returns singleton
func GetInstance() *Client {
	once.Do(func() {
		instance = &Client{
			proAPIKey: apiKey,
		}
	})
	return instance
}

// GetInstanceWithKey returns singleton
func GetInstanceWithKey(key string) *Client {
	once.Do(func() {
		instance = &Client{
			proAPIKey: key,
		}
	})
	return instance
}

// makeReq HTTP request helper
func (s *Client) makeReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-CMC_PRO_API_KEY", s.proAPIKey)
	return util.DoReq(req)
}

// Info returns all static metadata for one or more cryptocurrencies
// including name, symbol, logo, and its various registered URLs.
func (s *Client) Info(options *types.InfoOptions) (map[string]*types.Info, error) {
	var params []string
	if options == nil {
		options = new(types.InfoOptions)
	}
	if options.ID != "" {
		params = append(params, fmt.Sprintf("id=%s", options.ID))
	}
	if options.Symbol != "" {
		params = append(params, fmt.Sprintf("symbol=%s", options.Symbol))
	}

	url := fmt.Sprintf("%s/cryptocurrency/info?%s", baseURL, strings.Join(params, "&"))

	body, err := s.makeReq(url)
	resp := new(types.Response)
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	var result = make(map[string]*types.Info)
	ifcs, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, ErrCouldNotCast
	}

	for k, v := range ifcs {
		info := new(types.Info)
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(b, info); err != nil {
			return nil, err
		}
		result[k] = info
	}

	return result, nil
}

// ListingsLatest gets a paginated list of all cryptocurrencies with latest market data.
// You can configure this call to sort by market cap or another market ranking field.
// Use the "convert" option to return market values in multiple fiat and cryptocurrency conversions in the same call.
func (s *Client) ListingsLatest(options *types.ListingsLatestOptions) ([]*types.Listing, error) {
	var params []string
	if options == nil {
		options = new(types.ListingsLatestOptions)
	}
	if options.Start != 0 {
		params = append(params, fmt.Sprintf("start=%v", options.Start))
	}
	if options.Limit != 0 {
		params = append(params, fmt.Sprintf("limit=%v", options.Limit))
	}
	if options.Convert != "" {
		params = append(params, fmt.Sprintf("convert=%s", options.Convert))
	}
	if options.Sort != "" {
		params = append(params, fmt.Sprintf("sort=%s", options.Sort))
	}

	url := fmt.Sprintf("%s/cryptocurrency/listings/latest?%s", baseURL, strings.Join(params, "&"))

	body, err := s.makeReq(url)
	resp := new(types.Response)
	if err = json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}

	var listings []*types.Listing
	ifcs, ok := resp.Data.([]interface{})
	if !ok {
		return nil, ErrCouldNotCast
	}

	for i := range ifcs {
		ifc := ifcs[i]
		listing := new(types.Listing)
		b, err := json.Marshal(ifc)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(b, listing)
		if err != nil {
			return nil, err
		}
		listings = append(listings, listing)
	}

	return listings, nil
}
