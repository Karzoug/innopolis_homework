package main

import (
	"fmt"
	"maps"
	"sync"
)

func RunProcessor(wg *sync.WaitGroup, prices []*mmap[string, float64]) {
	go func() {
		defer wg.Done()
		for _, price := range prices {
			price.Map(func(key string, value float64) float64 {
				return value + 1
			})
			// for key, value := range price. {
			// 	price[key] = value + 1
			// }
			fmt.Println(price)
		}
	}()
}

func RunWriter() <-chan map[string]float64 {
	var prices = make(chan map[string]float64)
	go func() {
		var currentPrice = map[string]float64{
			"inst1": 1.1,
			"inst2": 2.1,
			"inst3": 3.1,
			"inst4": 4.1,
		}
		for i := 1; i < 5; i++ {
			for key, value := range currentPrice {
				currentPrice[key] = value + 1
			}
			prices <- maps.Clone(currentPrice)
			//time.Sleep(time.Second)
		}
		close(prices)
	}()
	return prices
}
func main() {
	p := RunWriter()
	var prices []*mmap[string, float64] //[string]float64

	for price := range p {
		prices = append(prices, NewMmap[string, float64](price))
	}

	for _, price := range prices {
		fmt.Println(price)
	}

	fmt.Println()

	wg := &sync.WaitGroup{}
	wg.Add(3)
	RunProcessor(wg, prices)
	RunProcessor(wg, prices)
	RunProcessor(wg, prices)
	wg.Wait()
}
