package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"sync"
	"time"
)

func main() {
	var price1, price2, price3 float64

	start := time.Now()

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		price1 = fetcher.FetchPrice1()
	}()

	go func() {
		defer wg.Done()
		price2 = fetcher.FetchPrice2()
	}()

	go func() {
		defer wg.Done()
		price3 = fetcher.FetchPrice3()
	}()

	wg.Wait()

	fmt.Printf("Site 1 - R$ %.2f \n", price1)
	fmt.Printf("Site 2 - R$ %.2f \n", price2)
	fmt.Printf("Site 3 - R$ %.2f \n", price3)

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
