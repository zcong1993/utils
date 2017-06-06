package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

// GetJson make a get request use http.Get
func GetJson(url string, v interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return err
}

// GetJsonWithHeaders make a http get request with custom headers
func GetJsonWithHeaders(url string, v interface{}, headers map[string]string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		return err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return err
}

// SliceIndex return the index of slice item which match the condition first
//
// will return -1 when not match all
func SliceIndex(limit int, condition func(num int) bool) int {
	for i := 0; i < limit; i++ {
		if condition(i) {
			return i
		}
	}
	return -1
}
