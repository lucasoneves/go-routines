package fetcher

import (
	"math/rand"
	"sync"
	"time"
)

func FetchPrices(priceChannel chan<- float64) {

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		priceChannel <- FetchPrice1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPrice2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPrice3()
	}()

	wg.Wait()
	close(priceChannel)
}

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

func FetchFromMultipleSite(priceChannel chan<- float64) {
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}

	for _, price := range prices {
		priceChannel <- price
	}

}
