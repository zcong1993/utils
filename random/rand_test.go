package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandRangeInt(t *testing.T) {
	type fixture [][2]int
	f := fixture{[2]int{1, 10}, [2]int{1, 5}, [2]int{-10, 10}}
	for _, v := range f {
		res := RandRangeInt(v[0], v[1])
		assert.True(t, res >= v[0])
		assert.True(t, res < v[1])
	}
}

func TestRandRangeFloat64(t *testing.T) {
	type fixture [][2]float64
	f := fixture{[2]float64{1, 10}, [2]float64{1, 5}, [2]float64{-10, 10}}
	for _, v := range f {
		res := RandRangeFloat64(v[0], v[1])
		assert.True(t, res >= v[0])
		assert.True(t, res < v[1])
	}
}
