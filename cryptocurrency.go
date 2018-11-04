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
func (s *Client) CryptoInfo(options *types.Options) (map[string]*types.CryptoInfo, error) {
	url := fmt.Sprintf("%s/cryptocurrency/info?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	resp, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var results = make(map[string]*types.CryptoInfo)
	data, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, ErrCouldNotCast
	}

	for k, v := range data {
		result := new(types.CryptoInfo)
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

// CryptoMap returns a paginated list of all cryptocurrencies by CoinMarketCap ID
// arg: symbol, start, limit, listing_status
// src: https://pro-api.coinmarketcap.com/v1/cryptocurrency/map
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1CryptocurrencyMap
func (s *Client) CryptoMap(options *types.Options) ([]*types.CryptoMap, error) {
	url := fmt.Sprintf("%s/cryptocurrency/map?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	resp, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var results []*types.CryptoMap
	data, ok := resp.Data.([]interface{})
	if !ok {
		return nil, ErrCouldNotCast
	}

	for i := range data {
		result := new(types.CryptoMap)
		b, err := json.Marshal(data[i])
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(b, result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

// CryptoListingsLatest gets a paginated list of all cryptocurrencies with latest market data.
// arg: start, limit, convert, sort, sort_dir, cryptocurrency_type
// src: https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest
// doc: https://pro.coinmarketcap.com/api/v1#operation/getV1ExchangeListingsLatest
func (s *Client) CryptoListingsLatest(options *types.Options) ([]*types.CryptoListing, error) {
	url := fmt.Sprintf("%s/cryptocurrency/listings/latest?%s", baseURL, strings.Join(util.ParseOptions(options), "&"))

	resp, err := s.getResponse(url)
	if err != nil {
		return nil, err
	}

	var results []*types.CryptoListing
	data, ok := resp.Data.([]interface{})
	if !ok {
		return nil, ErrCouldNotCast
	}

	for i := range data {
		result := new(types.CryptoListing)
		b, err := json.Marshal(data[i])
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(b, result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
