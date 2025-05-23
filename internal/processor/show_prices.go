package processor

import "fmt"

func ShowPriceAverage(priceChannel chan float64) {
	var totalPrice float64
	countPrices := 0
	for price := range priceChannel {
		totalPrice += price
		countPrices++
		averagePrice := totalPrice / float64(countPrices)
		fmt.Printf("Preço recebido: R$ %.2f | Preço médio: R$ %.2f\n", price, averagePrice)
	}
}
