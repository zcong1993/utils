package utils_test

import (
	"fmt"
	"github.com/zcong1993/utils"
)

// Example (simple) make a simple http get request for json data
func ExampleGetJSON() {
	// define a response data type
	type Resp struct {
		Status int    `decoder:"status"`
		Msg    string `decoder:"msg"`
	}
	var res Resp
	// this api is always return json data {"status":200, "msg": "hello world}
	err := utils.GetJSON("http://zcong-hello.getsandbox.com/hello", &res)
	// check error
	if err != nil {
		panic(err)
	}
	// res is response json data
	fmt.Printf("%v", res)
	//Output: {200 hello world}
}

// Example (custom headers) can add some custom headers when request
func ExampleGetJSONWithHeaders() {
	// define a response data type
	type Header struct {
		Name  string `decoder:"name"`
		Value string `decoder:"value"`
	}
	type Headers struct {
		Headers []Header `decoder:"headers"`
	}
	// define some custom headers
	customHeaders := map[string]string{"Foo": "bar"}
	var res Headers
	// this api return all the request headers
	err := utils.GetJSONWithHeaders("http://zcong-hello.getsandbox.com/header", &res, customHeaders)
	if err != nil {
		panic(err)
	}
	// custom headers should in response data
	index := utils.SliceIndex(len(res.Headers), func(num int) bool {
		return res.Headers[num] == Header{"Foo", "bar"}
	})
	fmt.Printf("%v", res.Headers[index])
	//Output: {Foo bar}
}
