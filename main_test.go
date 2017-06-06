package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetJson(t *testing.T) {
	type Resp struct {
		Status int    `decoder:"status"`
		Msg    string `decoder:"msg"`
	}
	var res Resp
	err := GetJson("http://zcong-hello.getsandbox.com/hello", &res)
	assert.Nil(t, err)
	assert.Equal(t, res, Resp{200, "hello world"})
}

func TestGetJsonWithHeaders(t *testing.T) {
	type Header struct {
		Name  string `decoder:"name"`
		Value string `decoder:"value"`
	}
	type Headers struct {
		Headers []Header `decoder:"headers"`
	}
	var res Headers
	err := GetJsonWithHeaders("http://zcong-hello.getsandbox.com/header", &res, map[string]string{"Content-Type": "application/json", "Foo": "bar"})
	assert.Nil(t, err)
	headers := res.Headers
	index1 := SliceIndex(len(headers), func(num int) bool {
		return headers[num] == Header{"Content-Type", "application/json"}
	})
	index2 := SliceIndex(len(headers), func(num int) bool {
		return headers[num] == Header{"Foo", "bar"}
	})
	assert.True(t, index1 > -1)
	assert.True(t, index2 > -1)
}

func TestSliceIndex(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	condition := func(i int) bool {
		return arr[i] == 3
	}
	index := SliceIndex(len(arr), condition)
	noExists := SliceIndex(len(arr), func (i int) bool {
		return arr[i] == 10
	})
	assert.Equal(t, index, 2)
	assert.Equal(t, noExists, -1)
}
