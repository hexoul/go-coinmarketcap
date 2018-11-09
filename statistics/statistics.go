// Package statistics gathers data and makes history
// It can be useful when your API plan is under standard, not authorized to historical data
package statistics

import (
	"fmt"
	"io"
	"os"
	"strings"

	kucoin "github.com/eeonevision/kucoin-go"
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
	timestampFormat := "02-01-2006 15:04:05"
	logger.Formatter = &log.JSONFormatter{
		TimestampFormat: timestampFormat,
	}
	if f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err == nil {
		logger.Out = io.MultiWriter(f, os.Stdout)
	} else {
		fmt.Print("Failed to open log file: you can miss important log")
		logger.Out = os.Stdout
	}
	logger.SetLevel(log.InfoLevel)
}

// GatherCryptoQuote records crypto quote
func GatherCryptoQuote(options *types.Options, job *gocron.Job) {
	job.Do(taskGatherCryptoQuote, options)
}

func taskGatherCryptoQuote(options *types.Options) {
	if data, err := coinmarketcap.GetInstance().CryptoMarketQuotesLatest(options); err == nil {
		for _, v := range data.CryptoMarket {
			logger.WithFields(log.Fields{
				"symbol": v.Symbol,
				"quote":  v.Quote,
			}).Info("GatherCryptoQuote")
		}
	}
}

// GatherExchangeMarketPairs records exchange quotes
func GatherExchangeMarketPairs(options *types.Options, targetSymbol string, job *gocron.Job) {
	job.Do(taskGatherExchangeMarketPairs, options, targetSymbol)
}

func taskGatherExchangeMarketPairs(options *types.Options, targetSymbol string) {
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

// GatherTokenMetric records the number of token holders and transactions
func GatherTokenMetric(symbol, addr string, job *gocron.Job) {
	job.Do(taskGatherTokenMetric, symbol, addr)
}

// symbol: Token symbol for log
// addr: Token contract address
// url: Target url to scraper
// c: Scraper
func taskGatherTokenMetric(symbol, addr string) {
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

// GatherKucoinBalance records a balance of Kucoin account
func GatherKucoinBalance(k *kucoin.Kucoin, symbol string, job *gocron.Job) {
	job.Do(taskGatherKucoinBalance, k, symbol)
}

func taskGatherKucoinBalance(k *kucoin.Kucoin, symbol string) {
	if bal, err := k.GetCoinBalance(symbol); err == nil {
		logger.WithFields(log.Fields{
			"symbol":  symbol,
			"balance": bal.Balance,
			"freeze":  bal.FreezeBalance,
		}).Info("GatherKucoinBalance")
	}
}
