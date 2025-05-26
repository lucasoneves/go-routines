package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/processor"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPriceAverage(priceChannel, done)

	<-done

	fmt.Printf("\nTempo total: %s \n", time.Since(start))
}
