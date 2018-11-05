// Package statistics gathers data and makes history
package statistics

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/jasonlvhit/gocron"
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
