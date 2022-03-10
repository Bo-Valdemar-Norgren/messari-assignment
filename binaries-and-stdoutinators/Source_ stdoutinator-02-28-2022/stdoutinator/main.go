package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"stdoutinator/models"
	"time"
)

const (
	tradeCount = 10000000 // ten million
)

func main() {
	// set random seed in a silly and fun manner
	rand.Seed(time.Now().Unix() * (time.Now().UnixMilli() % int64(time.Now().Second())))

	// create between 6,000 and 11,999 markets at random
	marketCount := (rand.Uint64() % 6000) + 6000

	beginTime := time.Now()

	fmt.Println("BEGIN")
	sendTrades(tradeCount, int(marketCount))
	fmt.Println("END")

	// report to the user the statistics for this run
	fmt.Println(fmt.Sprintf("Trade Count:  %d", tradeCount))
	fmt.Println(fmt.Sprintf("Market Count: %d\n", marketCount))
	fmt.Println(fmt.Sprintf("Duration of send operation: %s", time.Now().Sub(beginTime).String()))
}

// sendTrades prints random trades as JSON until the passed number have been printed.
func sendTrades(count int, marketCount int) {
	currentMarketID := 1
	countElapsed := 0

	for {
		if countElapsed > count {
			return
		}

		isBuy := rand.Uint64()%5 != 0 // the intention is that isBuy hovers around 20% in all cases

		t := models.Trade{
			ID:     countElapsed + 1,
			Market: currentMarketID,
			Price:  float64(currentMarketID%52) + rand.Float64(),
			Volume: rand.Float64() * 5000.00,
			IsBuy:  isBuy,
		}

		b, err := json.Marshal(t)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(b))

		countElapsed++
		if currentMarketID >= marketCount {
			currentMarketID = 1
		} else {
			currentMarketID++
		}
	}
}
