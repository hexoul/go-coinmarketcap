package types

// Options for request
type Options struct {
	ID         string
	Symbol     string
	Slug       string
	Start      int // >= 1
	Limit      int
	Convert    string // "USD", ...
	Sort       string // <- SortOptions
	SortDir    string // "asc", "desc"
	CryptoType string // "all", "coins", "tokens"
	MarketType string // "all", "fees", "no_fees"
}

// SortOptions sort options
var SortOptions sortOptions

type sortOptions struct {
	Name              string
	Symbol            string
	DateAdded         string
	MarketCap         string
	Price             string
	CirculatingSupply string
	TotalSupply       string
	MaxSupply         string
	NumMarketPairs    string
	Volume24H         string
	PercentChange1H   string
	PercentChange24H  string
	PercentChange7D   string
}

func init() {
	SortOptions = sortOptions{
		Name:              "name",
		Symbol:            "symbol",
		DateAdded:         "date_added",
		MarketCap:         "market_cap",
		Price:             "price",
		CirculatingSupply: "circulating_supply",
		TotalSupply:       "total_supply",
		MaxSupply:         "max_supply",
		NumMarketPairs:    "num_market_pairs",
		Volume24H:         "volume_24h",
		PercentChange1H:   "percent_change_1h",
		PercentChange24H:  "percent_change_24h",
		PercentChange7D:   "percent_change_7d",
	}
}
