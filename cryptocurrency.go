package coinmarketcap

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hexoul/go-coinmarketcap/types"
	"github.com/hexoul/go-coinmarketcap/util"
)

// CryptoInfo returns all static metadata for one or more cryptocurrencies
// arg: id, symbol
// src: https://pro-api.coinmarketcap.com/v1/cryptocurrency/info
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1CryptocurrencyInfo
func (s *Client) CryptoInfo(options *types.Options) (*types.CryptoInfoMap, error) {
	url := fmt.Sprintf("%s/cryptocurrency/info?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	_, body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	result := new(types.CryptoInfoMap)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CryptoMap returns a paginated list of all cryptocurrencies by CoinMarketCap ID
// arg: symbol, start, limit, listing_status
// src: https://pro-api.coinmarketcap.com/v1/cryptocurrency/map
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1CryptocurrencyMap
func (s *Client) CryptoMap(options *types.Options) (*types.CryptoMapList, error) {
	url := fmt.Sprintf("%s/cryptocurrency/map?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	_, body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	result := new(types.CryptoMapList)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CryptoListingsLatest gets a paginated list of all cryptocurrencies with latest market data.
// arg: start, limit, convert, sort, sort_dir, cryptocurrency_type
// src: https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1CryptocurrencyListingsLatest
func (s *Client) CryptoListingsLatest(options *types.Options) (*types.CryptoMarketList, error) {
	url := fmt.Sprintf("%s/cryptocurrency/listings/latest?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	_, body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	result := new(types.CryptoMarketList)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CryptoMarketPairsLatest lists all market pairs for the specified cryptocurrency with associated stats.
// arg: id, symbol, start, limit, convert
// src: https://pro-api.coinmarketcap.com/v1/cryptocurrency/market-pairs/latest
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1CryptocurrencyMarketpairsLatest
func (s *Client) CryptoMarketPairsLatest(options *types.Options) (*types.MarketPairs, error) {
	url := fmt.Sprintf("%s/cryptocurrency/market-pairs/latest?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

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

// CryptoMarketQuotesLatest gets the latest market quote for 1 or more cryptocurrencies.
// arg: id, symbol, convert
// src: https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1CryptocurrencyQuotesLatest
func (s *Client) CryptoMarketQuotesLatest(options *types.Options) (*types.CryptoMarketMap, error) {
	url := fmt.Sprintf("%s/cryptocurrency/quotes/latest?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	_, body, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	result := new(types.CryptoMarketMap)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}
