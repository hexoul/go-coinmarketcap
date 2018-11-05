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
	CryptoInfo(options *types.Options) (*types.CryptoInfoMap, error)
	CryptoMap(options *types.Options) (*types.CryptoMapList, error)
	CryptoListingsLatest(options *types.Options) (*types.CryptoMarketList, error)
	CryptoMarketPairsLatest(options *types.Options) (*types.MarketPairs, error)
	CryptoMarketQuotesLatest(options *types.Options) (*types.CryptoMarketMap, error)

	ExchangeInfo(options *types.Options) (*types.ExchangeInfoMap, error)
	ExchangeMap(options *types.Options) (*types.ExchangeMapList, error)
	ExchangeListingsLatest(options *types.Options) (*types.ExchangeMarketList, error)
	ExchangeMarketPairsLatest(options *types.Options) (*types.MarketPairs, error)
	ExchangeMarketQuotesLatest(options *types.Options) (*types.ExchangeMarketQuotes, error)
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
	baseURL = "https://pro-api.coinmarketcap.com/v1"
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
		if apiKey == "" {
			panic("API KEY REQUIRED")
		}
		instance = &Client{
			proAPIKey: apiKey,
		}
	})
	return instance
}

// GetInstanceWithKey returns singleton
func GetInstanceWithKey(key string) *Client {
	once.Do(func() {
		if key == "" {
			panic("API KEY REQUIRED")
		}
		instance = &Client{
			proAPIKey: key,
		}
	})
	return instance
}

func (s *Client) getResponse(url string) (*types.Response, []byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("X-CMC_PRO_API_KEY", s.proAPIKey)
	body, err := util.DoReq(req)
	if err != nil {
		return nil, nil, err
	}
	resp := new(types.Response)
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, nil, err
	}
	if resp.Status.ErrorCode != 0 {
		return nil, nil, fmt.Errorf("[%d] %s", resp.Status.ErrorCode, *resp.Status.ErrorMessage)
	}
	return resp, body, nil
}
