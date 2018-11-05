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
func (s *Client) ExchangeInfo(options *types.Options) (*types.ExchangeInfoMap, error) {
	url := fmt.Sprintf("%s/exchange/info?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	_, body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	result := new(types.ExchangeInfoMap)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExchangeMap returns a paginated list of all cryptocurrency exchanges by CoinMarketCap ID.
// arg: slug, start, limit, listing_status
// src: https://pro-api.coinmarketcap.com/v1/exchange/map
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1ExchangeMap
func (s *Client) ExchangeMap(options *types.Options) (*types.ExchangeMapList, error) {
	url := fmt.Sprintf("%s/exchange/map?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	_, body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	result := new(types.ExchangeMapList)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExchangeListingsLatest gets a paginated list of all cryptocurrency exchanges including the latest aggregate market data for each exchange.
// arg: start, limit, sort, sort_dir, market_type, convert
// src: https://pro-api.coinmarketcap.com/v1/exchange/listings/latest
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1ExchangeListingsLatest
func (s *Client) ExchangeListingsLatest(options *types.Options) (*types.ExchangeMarketList, error) {
	url := fmt.Sprintf("%s/exchange/listings/latest?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	_, body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var result = new(types.ExchangeMarketList)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExchangeMarketPairsLatest get a list of active market pairs for an exchange.
// arg: id, slug, start, limit, convert
// src: https://pro-api.coinmarketcap.com/v1/exchange/market-pairs/latest
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1ExchangeMarketpairsLatest
func (s *Client) ExchangeMarketPairsLatest(options *types.Options) (*types.MarketPairs, error) {
	url := fmt.Sprintf("%s/exchange/market-pairs/latest?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	resp, _, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var result = new(types.MarketPairs)
	b, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExchangeMarketQuotesLatest gets the latest aggregate market data for 1 or more exchanges.
// arg: id, slug, convert
// src: https://pro-api.coinmarketcap.com/v1/exchange/quotes/latest
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1ExchangeQuotesLatest
func (s *Client) ExchangeMarketQuotesLatest(options *types.Options) (*types.ExchangeMarketQuotes, error) {
	url := fmt.Sprintf("%s/exchange/quotes/latest?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	_, body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var result = new(types.ExchangeMarketQuotes)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}
