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

	var result = make(map[string]*types.CryptoInfo)
	ifcs, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, ErrCouldNotCast
	}

	for k, v := range ifcs {
		info := new(types.CryptoInfo)
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

	var listings []*types.CryptoListing
	ifcs, ok := resp.Data.([]interface{})
	if !ok {
		return nil, ErrCouldNotCast
	}

	for i := range ifcs {
		ifc := ifcs[i]
		listing := new(types.CryptoListing)
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
