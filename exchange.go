package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hexoul/go-coinmarketcap/types"
	"github.com/hexoul/go-coinmarketcap/util"
)

// ExchangeInfo returns all static metadata for one or more exchanges including logo and homepage URL.
// arg: id, slug
// src: https://pro-api.coinmarketcap.com/v1/exchange/info
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1ExchangeInfo
func (s *Client) ExchangeInfo(options *types.Options) (map[string]*types.ExchangeInfo, error) {
	url := fmt.Sprintf("%s/exchange/info?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	resp, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var results = make(map[string]*types.ExchangeInfo)
	data, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, ErrCouldNotCast
	}

	for k, v := range data {
		result := new(types.ExchangeInfo)
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(b, result); err != nil {
			return nil, err
		}
		results[k] = result
	}

	return results, nil
}

// ExchangeMarketPairsLatest get a list of active market pairs for an exchange.
// arg: id, slug, start, limit, convert
// src: https://pro-api.coinmarketcap.com/v1/exchange/market-pairs/latest
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1ExchangeMarketpairsLatest
func (s *Client) ExchangeMarketPairsLatest(options *types.Options) (*types.ExchangeMarketPairs, error) {
	url := fmt.Sprintf("%s/exchange/info?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	resp, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var result = new(types.ExchangeMarketPairs)
	data, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, ErrCouldNotCast
	}
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, result); err != nil {
		return nil, err
	}
	return result, nil
}
