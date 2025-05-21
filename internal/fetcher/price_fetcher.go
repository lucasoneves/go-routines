package fetcher

import (
	"math/rand"
	"time"
)

func FetchPrice1() float64 {
	time.Sleep(1 * time.Second)
	return rand.Float64() * 100
}

func FetchPrice2() float64 {
	time.Sleep(3 * time.Second)
	return rand.Float64() * 100
}

func FetchPrice3() float64 {
	time.Sleep(5 * time.Second)
	return rand.Float64() * 100
}
