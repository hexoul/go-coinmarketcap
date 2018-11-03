// Package util supports specific parsing
package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// ToInt helper for parsing strings to int
func ToInt(rawInt string) int {
	parsed, _ := strconv.Atoi(strings.Replace(strings.Replace(rawInt, "$", "", -1), ",", "", -1))
	return parsed
}

// ToFloat helper for parsing strings to float
func ToFloat(rawFloat string) float64 {
	parsed, _ := strconv.ParseFloat(strings.Replace(strings.Replace(strings.Replace(rawFloat, "$", "", -1), ",", "", -1), "%", "", -1), 64)
	return parsed
}

// DoReq HTTP client
func DoReq(req *http.Request) (body []byte, err error) {
	requestTimeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: requestTimeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else if resp != nil && resp.StatusCode != 200 {
		err = fmt.Errorf("%s", body)
	}
	return
}
