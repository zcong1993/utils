package random

import (
	"math/rand"
	"time"
)

// RandInt is helper generate random int func auto seed unix nano
func RandInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

// RandFloat64 is helper generate random float64 func auto seed unix nano
func RandFloat64() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
}

// RandRangeInt is helper func generate int with range
func RandRangeInt(min, max int) int {
	if min >= max {
		return 0
	}
	r := RandInt()
	return min + r%(max-min)
}

// RandRangeFloat64 is helper func generate float64 with range
func RandRangeFloat64(min, max float64) float64 {
	if min >= max {
		return 0
	}
	r := RandFloat64()
	return min + r*(max-min)
}
