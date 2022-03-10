package main

import (
	"bufio"
	"fmt"
	"os"
)


// Probably requires some tracker variables for efficient calculation of meanprice, VWAP, percentageBuyOrder
type Order struct {
	id int
	market int
	price float32
	volume float32
	is_buy bool
}


func main() {
	input := bufio.NewScanner(os.Stdin)
	metrics := make(map[int]map[string]float32)

	for input.Scan() {
		line := input.Text()

		// Ready to receive first order
		if line == "BEGIN" {
			continue
		// Last order was reached
		} else if line == "END" {
			break
		// Receieve order
		} else {
			var order Order
			err := json.Unmarshal([]byte(line), &order)
			if err != nil {
				fmt.Println("ERROR: Order could not be unmarshaled.")
			} else {
				//fmt.Println(order)
				processOrder(metrics, &order)
			}
		}
	}
	outputMetrics(metrics)
}


func processOrder(metrics map[int]map[string]float32, order *Order) {
	marketID := order.market

	// Market has previously been traded on, update metrics for it.
	if metrics[marketID] {
		updateMetrics(metrics, order)

	// First trade on market, initialize metrics for it.
	} else {
		initializeMetrics(metrics, order)
	}
}


func initializeMetrics(metrics map[int]map[string]float32, order *Order) {
	marketID := order.market
	price := order.price
	volume := order.volume
	isBuy := order.is_buy

	metrics[marketID] = make(map[string]float32)

	metrics[marketID]["totalVolume"] = volume
	metrics[marketID]["meanPrice"] = price
	metrics[marketID]["meanVolume"] = volume
	metrics[marketID]["VWAP"] = price

	var percentageBuy float32 = 0
	if isBuy {
		percentageBuy = 1
	}
	metrics[marketID]["percentageBuyOrder"] = percentageBuy
}


func updateMetrics(metrics map[int]map[string]float32, order *Order) {
	marketID := order.market
	price := order.price
	volume := order.volume
	isBuy := order.is_buy

	// Update total volume
	metrics[marketID]["totalVolume"] += volume

	// TODO: Find an efficient way of continously updating these metrics

	// Update mean price

	// Update mean volume

	// Update VWAP

	// Update percentage buy order
}


func outputMetrics (metrics map[int]map[string]float32) {
	for marketID, marketMetrics := range metrics {
		totalVolume := marketMetrics["totalVolume"]
		meanPrice := marketMetrics["meanPrice"]
		meanVolume := marketMetrics["meanVolume"]
		vWAP := marketMetrics["VWAP"]
		percentageBuyOrder := marketMetrics["percentageBuyOrder"]

		fmt.Println("{\"market\":%d, \"total_volume\":%g, \"mean_price\":%g,\"mean_volume\":%g, \"volume_weighted_average_price\":%g, \"percentage_buy\":%g,  }",
		 marketID, totalVolume, meanPrice, meanVolume, vWAP, percentageBuyOrder)
	} 
}
