package fetcher

import (
	"buscador/models"
	"math/rand"
	"sync"
	"time"
)

func FetchPrices(priceChannel chan<- models.ProductDetail) {

	var wg sync.WaitGroup
	wg.Add(4)

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

	go func() {
		defer wg.Done()
		FetchFromMultipleSite(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

func FetchPrice1() models.ProductDetail {
	time.Sleep(1 * time.Second)
	return models.ProductDetail{
		StorageName: "Magalu",
		Value:       rand.Float64() * 100,
		TimeStamp:   time.Now(),
	}
}

func FetchPrice2() models.ProductDetail {
	time.Sleep(3 * time.Second)
	return models.ProductDetail{
		StorageName: "Pontofrio",
		Value:       rand.Float64() * 100,
		TimeStamp:   time.Now(),
	}
}

func FetchPrice3() models.ProductDetail {
	time.Sleep(5 * time.Second)
	return models.ProductDetail{
		StorageName: "Amazon",
		Value:       rand.Float64() * 100,
		TimeStamp:   time.Now(),
	}
}

func FetchFromMultipleSite(priceChannel chan<- models.ProductDetail) {
	time.Sleep(2 * time.Second)
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}

	for _, price := range prices {
		priceChannel <- models.ProductDetail{
			StorageName: "MultiSite",
			Value:       price,
			TimeStamp:   time.Now(),
		}
	}

}
