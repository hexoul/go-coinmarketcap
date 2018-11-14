// Package statistics gathers data and makes history
// It can be useful when your API plan is under standard, not authorized to historical data
package statistics

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"

	coinmarketcap "github.com/hexoul/go-coinmarketcap"
	"github.com/hexoul/go-coinmarketcap/types"
	"github.com/hexoul/go-coinmarketcap/util"
)

var (
	logger *log.Logger
)

func init() {
	logPath := "./report.log"
	for _, val := range os.Args {
		arg := strings.Split(val, "=")
		if len(arg) < 2 {
			continue
		} else if arg[0] == "-logpath" {
			logPath = arg[1]
		}
	}

	// Initialize logger
	logger = log.New()

	// Default configuration
	logger.Formatter = &log.JSONFormatter{
		TimestampFormat: time.RFC3339,
	}
	if f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err == nil {
		logger.Out = io.MultiWriter(f, os.Stdout)
	} else {
		fmt.Print("Failed to open log file: you can miss important log")
		logger.Out = os.Stdout
	}
	logger.SetLevel(log.InfoLevel)
}

// GatherCryptoQuote gathers crypto quote
func GatherCryptoQuote(options *types.Options, job *gocron.Job) {
	job.Do(TaskGatherCryptoQuote, options)
}

// TaskGatherCryptoQuote records crypto quote
func TaskGatherCryptoQuote(options *types.Options) {
	if data, err := coinmarketcap.GetInstance().CryptoMarketQuotesLatest(options); err == nil {
		for _, v := range data.CryptoMarket {
			logger.WithFields(log.Fields{
				"symbol": v.Symbol,
				"quote":  v.Quote,
			}).Info("GatherCryptoQuote")
		}
	}
}

// GatherOhlcv gathers crypto quote
func GatherOhlcv(options *types.Options, job *gocron.Job) {
	job.Do(TaskGatherOhlcv, options)
}

// TaskGatherOhlcv records crypto quote
func TaskGatherOhlcv(options *types.Options) {
	start := time.Now().AddDate(0, 0, -2)
	end := time.Now().AddDate(0, 0, -1)
	options.TimeStart = util.ISODate(start)
	options.TimeEnd = util.ISODate(end)

	if data, err := coinmarketcap.GetInstance().CryptoOhlcvHistorical(options); err == nil {
		for _, v := range data.Ohlcv {
			for _, q := range v.Quotes {
				logger.WithFields(log.Fields{
					"symbol": data.Symbol,
					"quote":  q,
					"ctime":  v.TimeClose,
				}).Info("GatherOhlcv")
			}
		}
	}
}

// GatherExchangeMarketPairs gathers exchange quotes
func GatherExchangeMarketPairs(options *types.Options, targetSymbol string, job *gocron.Job) {
	job.Do(TaskGatherExchangeMarketPairs, options, targetSymbol)
}

// TaskGatherExchangeMarketPairs records exchange quotes
func TaskGatherExchangeMarketPairs(options *types.Options, targetSymbol string) {
	if data, err := coinmarketcap.GetInstance().ExchangeMarketPairsLatest(options); err == nil {
		for _, pair := range data.MarketPair {
			if strings.Contains(pair.MarketPair, targetSymbol) {
				logger.WithFields(log.Fields{
					"symbol": targetSymbol,
					"market": data.Slug,
					"pair":   pair.MarketPair,
					"quote":  pair.Quote,
				}).Info("GatherExchangeMarketPairs")
			}
		}
	}
}

// GatherTokenMetric gathers the number of token holders and transactions
func GatherTokenMetric(symbol, addr string, job *gocron.Job) {
	job.Do(TaskGatherTokenMetric, symbol, addr)
}

// TaskGatherTokenMetric records the number of token holders and transactions
// symbol: Token symbol for log
// addr: Token contract address
func TaskGatherTokenMetric(symbol, addr string) {
	var holders, txns, transfers string

	// Target urls
	url1 := "https://etherscan.io/token/" + addr
	url2 := "https://etherscan.io/address/" + addr

	// Invoke chromedp
	util.InvokeChromedp(
		url1,
		"#ContentPlaceHolder1_divSummary td span#totaltxns",
		5,
		&transfers,
	)

	// Initialize collector
	c1 := colly.NewCollector()
	c2 := colly.NewCollector()

	c1.OnHTML("#ContentPlaceHolder1_tr_tokenHolders", func(e *colly.HTMLElement) {
		holders = strings.Split(e.ChildText("td:nth-of-type(2)"), " ")[0]
	})
	c2.OnHTML(".table tbody tr td span", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "txns") {
			txns = strings.Split(e.Text, " ")[0]
		}
	})

	c1.Visit(url1)
	c2.Visit(url2)

	logger.WithFields(log.Fields{
		"symbol":    symbol,
		"holders":   holders,
		"transfers": transfers,
		"txns":      txns,
	}).Info("GatherTokenMetric")
}
