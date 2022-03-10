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
	buyOrders           int
	sellOrders          int
	percentageBuyOrders float32
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
		// Ready to receive first order
		if line == "BEGIN" {
			continue
		// Last order was reached
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
	}

	metric := metrics[order.Market]

	if order.IsBuy {
		metric.buyOrders = 1
		metric.percentageBuyOrders = 1.0
	} else {
		metric.sellOrders = 1
		metric.percentageBuyOrders = 0.0
	}
}


func updateMarket(metrics map[int]*Metrics, order *Order) {
	metric := metrics[order.Market]

	// Update total volume
	metric.totalVolume += order.Volume

	// TODO: Find an efficient way of continously updating these metrics

	// Update mean price

	// Update mean volume

	// Update VWAP

	// Update percentage buy orders
}


func outputMetrics(metrics map[int]*Metrics) {
	for marketID, metrics := range metrics {
		fmt.Printf("{\"market\":%d, \"total_volume\":%g, \"mean_price\":%g,\"mean_volume\":%g, \"volume_weighted_average_price\":%g, \"percentage_buy\":%g}",
		 marketID, metrics.totalVolume, metrics.meanPrice, metrics.meanVolume, metrics.VWAP, metrics.percentageBuyOrders)
	} 
}
