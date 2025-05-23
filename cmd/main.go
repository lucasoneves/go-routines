package main

import (
	"buscador/internal/fetcher"
	"buscador/internal/processor"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	priceChannel := make(chan float64)
	var consumerWg sync.WaitGroup
	consumerWg.Add(1)

	go func() {
		defer consumerWg.Done() // Garante que consumerWg.Done() seja chamado quando o consumidor terminar
		processor.ShowPriceAverage(priceChannel)
	}()

	go fetcher.FetchPrices(priceChannel)

	// A função main agora espera que a go routine consumidora termine
	// de processar todos os preços e que o canal seja fechado.
	consumerWg.Wait()

	fmt.Printf("\nTempo total: %s \n", time.Since(start))
}
