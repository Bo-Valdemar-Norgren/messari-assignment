package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
)

type Metrics struct {
	totalVolume         float32
	meanPrice           float32
	meanVolume          float32
	VWAP                float32
	buyOrders           float32
	sellOrders          float32
	percentageBuyOrders float32 
	orderCounter        float32
	totalPriceVolume    float32
}

type Order struct {
	ID     int     `json:"id"`
	Market int     `json:"market"`
	Price  float32 `json:"price"`
	Volume float32 `json:"volume"`
	IsBuy  bool    `json:"is_buy"`
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	metrics := make(map[int]*Metrics)

	for input.Scan() {
		line := input.Text()

		if line == "BEGIN" {
			continue
		} else if line == "END" {
			break
		// Receive order
		} else {
			var order Order
			err := json.Unmarshal([]byte(line), &order)
			if err != nil {
				fmt.Println("ERROR: Order could not be unmarshaled.")
			} else {
				processOrder(metrics, &order)
			}
		}
	}
	outputMetrics(metrics)
}

func processOrder(metrics map[int]*Metrics, order *Order) {
	// Market has previously been traded on, update metrics for it.
	if _, exists := metrics[order.Market]; exists {
		updateMarket(metrics, order)
	// First trade on market, initialize metrics for it.
	} else {
		initializeMarket(metrics, order)
	}
}

func initializeMarket(metrics map[int]*Metrics, order *Order) {
	metrics[order.Market] = &Metrics{
		totalVolume: order.Volume,
		meanPrice: order.Price,
		meanVolume: order.Volume,
		VWAP: order.Price,
		percentageBuyOrders: 0,
		orderCounter: 1,
		totalPriceVolume: order.Price * order.Volume,
	}

	metric := metrics[order.Market]

	if order.IsBuy {
		metric.buyOrders = 1
		metric.percentageBuyOrders = 1
	} else {
		metric.sellOrders = 1
	}
}

func updateMarket(metrics map[int]*Metrics, order *Order) {
	metric := metrics[order.Market]
	metric.orderCounter += 1

	if order.IsBuy {
		metric.buyOrders += 1
	} else {
		metric.sellOrders += 1
	}

	// Update mean price
	metric.meanPrice = metric.meanPrice + (order.Price - metric.meanPrice) / metric.orderCounter

	// Update total volume
	metric.totalVolume += order.Volume

	// Update total price*volume
	metric.totalPriceVolume += order.Price * order.Volume

	// Update VWAP
	metric.VWAP = metric.totalPriceVolume / metric.totalVolume

	// Update mean volume
	metric.meanVolume = metric.totalVolume / metric.orderCounter

	// Update percentage buy orders
	metric.percentageBuyOrders = metric.buyOrders / (metric.buyOrders + metric.sellOrders)
}

func outputMetrics(metrics map[int]*Metrics) {
	for marketID, metric := range metrics {
		fmt.Printf("{\"market\":%d, \"total_volume\":%g, \"mean_price\":%g,\"mean_volume\":%g, \"volume_weighted_average_price\":%g, \"percentage_buy\":%g}\n",
		    marketID, metric.totalVolume, metric.meanPrice, metric.meanVolume, metric.VWAP, metric.percentageBuyOrders)
	} 
}
