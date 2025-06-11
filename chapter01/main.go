package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stock struct {
	Symbol string
	Price  float64
}

func fetchStockPrice(symbol string) Stock {
	delay := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(delay)

	price := 100 + rand.Float64()*50

	return Stock{
		Symbol: symbol,
		Price:  price,
	}
}

func main() {
	stocks := []string{
		"AAPL", "TSLA", "GOOGLE", "NAVER",
	}

	start := time.Now()

	var results []Stock
	for _, symbol := range stocks {
		go func(s string) {
			stock := fetchStockPrice(symbol)
			results = append(results, stock)
		}(symbol)
	}

	elapsedTime := time.Since(start)
	time.Sleep(4 * time.Second) // Goroutine 을 쉽게 기다리는법(좋지 않은 방법)

	fmt.Printf("총 소요 시간 :%v, 결과: %v\n", elapsedTime, results)
}
