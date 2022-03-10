package main

import (
	"bufio"
	"fmt"
	"os"
)


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
	totalOrders := 0

	// TODO: Add "structure" of following lines as the default structure for metrics
	// metrics[marketID]["totalVolume"]
	// metrics[marketID]["meanPrice"]
	// metrics[marketID]["meanVolume"]
	// metrics[marketID]["VWAP"]
	// metrics[marketID]["percentageBuyOrder"]

	for input.Scan() {
		line := input.Text()

		if line == "BEGIN" {
			continue
		} else if line == "END" {
			break
		} else {
			var order Order =
			parseOrderAndUpdateMetrics(metrics, &order)
		}
	}
	outputMetrics(metrics)
}


func parseOrderAndUpdateMetrics(metrics map[int]map[string]float32, order *Order) {
	marketID := order.market

	// If market has been traded on
	if metrics[marketID] {


	// If market has not yet been traded on
	} else {

	}
}


func initializeMarketMetrics(metrics map[int]map[string]float32, order *Order) {
	// Probably requires some tracker variables for meanprice, VWAP, percentageBuyOrder

	marketID := 
	price :=
	volume :=
	isBuy :=

	metrics[marketID] = make(map[string]float32)

	metrics[marketID]["totalVolume"] = volume
	metrics[marketID]["meanPrice"] = price
	metrics[marketID]["VWAP"] = price

	var percentageBuy float32 = 0
	if isBuy {
		percentageBuy = 1
	}
	metrics[marketID]["percentageBuyOrder"] = percentageBuy

}


func updateMarketMetrics(metrics map[int]map[string]float32, order *Order) {
	marketID := 
	price :=
	volume :=
	isBuy :=

	// Update total volume
	metrics[marketID]["totalVolume"] += volume

	// Update mean price

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
