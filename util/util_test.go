package util

import (
	"testing"

	"github.com/hexoul/go-coinmarketcap/types"
)

func TestParseOptions(t *testing.T) {
	if len(ParseOptions(nil)) != 0 {
		t.FailNow()
	}
	if len(ParseOptions(&types.Options{
		Symbol: "BTC",
	})) != 1 {
		t.FailNow()
	}
	if len(ParseOptions(&types.Options{
		Limit: 1,
		Sort:  types.SortOptions.Name,
	})) != 2 {
		t.FailNow()
	}
}
