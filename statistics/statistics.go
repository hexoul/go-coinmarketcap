// Package statistics gathers data and makes history
// It can be useful when your API plan is under standard, not authorized to historical data
package statistics

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/jasonlvhit/gocron"
	log "github.com/sirupsen/logrus"

	coinmarketcap "github.com/hexoul/go-coinmarketcap"
	"github.com/hexoul/go-coinmarketcap/types"
)

var (
	logger *log.Logger
)

func init() {
	// Initialize logger
	logger = log.New()

	// Default configuration
	timestampFormat := "02-01-2006 15:04:05"
	logger.Formatter = &log.JSONFormatter{
		TimestampFormat: timestampFormat,
	}
	if f, err := os.OpenFile("./report.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err == nil {
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

// GatherTokenMetric records the number of token holders and transactions
func GatherTokenMetric(symbol, addr string, job *gocron.Job) {
	job.Do(taskGatherTokenMetric, symbol, addr)
}

// symbol: Token symbol for log
// addr: Token contract address
// url: Target url to scraper
// c: Scraper
func taskGatherTokenMetric(symbol, addr string) {
	var holders, txns string
	// Initialize collector
	url1 := "https://etherscan.io/token/" + addr
	url2 := "https://etherscan.io/address/" + addr

	c1 := colly.NewCollector()
	c2 := colly.NewCollector()

	c1.OnHTML(".table tbody #ContentPlaceHolder1_tr_tokenHolders", func(e *colly.HTMLElement) {
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
		"symbol":  symbol,
		"holders": holders,
		"txns":    txns,
	}).Info("GatherTokenMetric")
}

func testLog() {
	logger.WithFields(log.Fields{
		"market":      "binance",
		"market_pair": "ETH/BTC",
	}).Info("TEST")
}

func testCron() {
	gocron.Every(1).Minute().Do(testLog)
	gocron.Every(2).Seconds().Do(testLog)
	gocron.Every(1).Day().At("09:35").Do(testLog)
	gocron.Start()
}
