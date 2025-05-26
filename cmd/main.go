package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/processor"
	"buscador/models"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan models.ProductDetail)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAverage(priceChannel, done)

	<-done

	fmt.Printf("\nTempo total: %s \n", time.Since(start))
}
