package types

// Options for request
type Options struct {
	ID         string `json:"id,omitempty"`
	Symbol     string `json:"symbol,omitempty"`
	Slug       string `json:"slug,omitempty"`
	Start      int    `json:"start,omitempty"` // >= 1
	Limit      int    `json:"limit,omitempty"`
	Convert    string `json:"convert,omitempty"`             // "USD", ...
	Sort       string `json:"sort,omitempty"`                // <- SortOptions
	SortDir    string `json:"sort_dir,omitempty"`            // "asc", "desc"
	CryptoType string `json:"cryptocurrency_type,omitempty"` // "all", "coins", "tokens"
	MarketType string `json:"market_type,omitempty"`         // "all", "fees", "no_fees"
	TimePeriod string `json:"time_period,omitempty"`         // "daily"
	TimeStart  string `json:"time_start,omitempty"`
	TimeEnd    string `json:"time_end,omitempty"`
	Interval   string `json:"interval,omitempty"` // "hourly" "daily" "weekly" "monthly" "yearly" "1d" "2d" "3d" "7d" "14d" "15d" "30d" "60d" "90d" "365d"
	Count      int    `json:"count,omitempty"`
}

type intervalOptions struct {
	Hourly  string
	Daily   string
	Weekly  string
	Montly  string
	Yearly  string
	Day1    string
	Days2   string
	Days3   string
	Days7   string
	Days14  string
	Days15  string
	Days30  string
	Days60  string
	Days90  string
	Days365 string
}

// IntervalOptions for interval
var IntervalOptions intervalOptions

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

// SortOptions for sorting
var SortOptions sortOptions

func init() {
	IntervalOptions = intervalOptions{
		Hourly:  "hourly",
		Daily:   "daily",
		Weekly:  "weekly",
		Montly:  "montly",
		Yearly:  "yearly",
		Day1:    "1d",
		Days2:   "2d",
		Days3:   "3d",
		Days7:   "7d",
		Days14:  "14d",
		Days15:  "15d",
		Days30:  "30d",
		Days60:  "60d",
		Days90:  "90d",
		Days365: "365d",
	}

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
