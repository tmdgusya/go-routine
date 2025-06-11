package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Stock struct {
	Symbol string
	Price  float64
	Time   time.Time
}

func fetchStockPrice(symbol string) Stock {
	delay := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(delay)

	price := 100 + rand.Float64()*50
	return Stock{
		Symbol: symbol,
		Price:  price,
		Time:   time.Now(),
	}
}

func stockPriceWithChannels() {
	stocks := []string{"APPLE", "TSLA", "GOOGLE", "NAVER"}

	channels := make([]chan Stock, len(stocks)) // buffer
	for i := range channels {
		channels[i] = make(chan Stock)
	}

	for i, symbol := range stocks {
		go func(i int, s string) {
			stock := fetchStockPrice(s)
			channels[i] <- stock
		}(i, symbol)
	}

	var results []Stock
	for _, ch := range channels {
		stock := <-ch
		results = append(results, stock)
	}

	fmt.Printf("수신 완료: %v\n", results)
}

func basicChannelDemo() {
	ch := make(chan string)
	go func() {
		fmt.Println("고루틴: 메세지 전송 중...")
		time.Sleep(2 * time.Second)
		ch <- "Hello from goroutine"
		fmt.Println("고루틴: 메세지 전송 완료")
	}()

	fmt.Println("메세지 수신 대기 중...")
	message := <-ch
	fmt.Printf("받은 메시지: %s\n", message)
}

func main() {
	basicChannelDemo()
	stockPriceWithChannels()
}
