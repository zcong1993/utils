package qs

import (
	"github.com/magiconair/properties/assert"
	"net/url"
	"testing"
)

func TestRawEncode(t *testing.T) {
	type fixture struct {
		in   map[string]string
		want string
	}
	fixtures := []fixture{
		{in: map[string]string{}, want: ""},
		{in: map[string]string{"a": "a", "b": "b"}, want: "a=a&b=b"},
		{in: map[string]string{"b": "b", "a": "a"}, want: "a=a&b=b"},
	}

	for _, f := range fixtures {
		vv := url.Values{}
		for k, v := range f.in {
			vv.Add(k, v)
		}

		got := RawEncode(vv)

		assert.Equal(t, got, f.want)
	}

	v := url.Values{}
	v.Add("a", "a1")
	v.Add("a", "a2")

	assert.Equal(t, RawEncode(v), "a=a1&a=a2")
}
