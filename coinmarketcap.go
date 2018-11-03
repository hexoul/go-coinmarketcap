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
	Info(options *types.Options) (map[string]*types.Info, error)
	ListingsLatest(options *types.Options) ([]*types.Listing, error)
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
// src: https://pro-api.coinmarketcap.com/v1/cryptocurrency/info
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1CryptocurrencyInfo
func (s *Client) Info(options *types.Options) (map[string]*types.Info, error) {
	url := fmt.Sprintf("%s/cryptocurrency/info?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	body, err := s.makeReq(url)
	if err != nil {
		return nil, err
	}
	resp := new(types.Response)
	if err := json.Unmarshal(body, &resp); err != nil {
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
		if err := json.Unmarshal(b, info); err != nil {
			return nil, err
		}
		result[k] = info
	}

	return result, nil
}

// ListingsLatest gets a paginated list of all cryptocurrencies with latest market data.
// You can configure this call to sort by market cap or another market ranking field.
// Use the "convert" option to return market values in multiple fiat and cryptocurrency conversions in the same call.
// src: https://pro-api.coinmarketcap.com/v1/exchange/listings/latest
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1ExchangeListingsLatest
func (s *Client) ListingsLatest(options *types.Options) ([]*types.Listing, error) {
	url := fmt.Sprintf("%s/cryptocurrency/listings/latest?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	body, err := s.makeReq(url)
	if err != nil {
		return nil, err
	}
	resp := new(types.Response)
	if err := json.Unmarshal(body, &resp); err != nil {
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
		if err := json.Unmarshal(b, listing); err != nil {
			return nil, err
		}
		listings = append(listings, listing)
	}

	return listings, nil
}
