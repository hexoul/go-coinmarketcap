package statistics

import (
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	testLog()
}

func TestCron(t *testing.T) {
	testCron()
	time.Sleep(20 * time.Second)
}
