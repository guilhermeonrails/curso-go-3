package fetcher

import (
	"concorrencia/internal/models"
	"math/rand"
	"sync"
	"time"
)

// Buscando preços de diferentes sites
func FetchPrices(priceChannel chan<- models.PriceDetail) {
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite3()
	}()

	go func() {
		defer wg.Done()
		FetchMultiplePriceFromSite4(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

func FetchPriceFromSite1() models.PriceDetail {
	time.Sleep(1 * time.Second)
	return models.PriceDetail{
		StoreName: "A",
		Value:     rand.Float64() * 100,
		TimeStamp: time.Now(),
	}
}

func FetchPriceFromSite2() models.PriceDetail {
	time.Sleep(3 * time.Second)
	return models.PriceDetail{
		StoreName: "B",
		Value:     rand.Float64() * 100,
		TimeStamp: time.Now(),
	}
}

func FetchPriceFromSite3() models.PriceDetail {
	time.Sleep(2 * time.Second)
	return models.PriceDetail{
		StoreName: "A",
		Value:     rand.Float64() * 100,
		TimeStamp: time.Now(),
	}
}

func FetchMultiplePriceFromSite4(priceChannel chan<- models.PriceDetail) {
	time.Sleep(6 * time.Second)
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}
	for _, price := range prices {
		priceChannel <- models.PriceDetail{
			StoreName: "D",
			Value:     price,
			TimeStamp: time.Now(),
		}
	}
}
