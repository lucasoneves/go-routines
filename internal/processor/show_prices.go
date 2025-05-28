package processor

import (
	"buscador/models"
	"fmt"
)

func ShowPriceAverage(priceChannel <-chan models.ProductDetail, done chan<- bool) {
	var totalPrice float64
	countPrices := 0
	for price := range priceChannel {
		totalPrice += price.Value
		countPrices++
		averagePrice := totalPrice / float64(countPrices)

		fmt.Printf("Data: %s | Preço de %s | R$ %.2f - Preço médio: R$ %.2f \n", price.TimeStamp.Format("02/01/2006 15:04"), price.StorageName, price.Value, averagePrice)
	}
	done <- true
}
