package main

import (
	"appmagic-entrance/analytics"
	"appmagic-entrance/parser"
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://raw.githubusercontent.com/CryptoRStar/GasPriceTestTask/main/gas_price.json", "URL to json file")
	outputFlag := flag.String("output", "result.json", "Path to store result")
	flag.Parse()

	url := string(*urlFlag)
	output := string(*outputFlag)

	fmt.Printf("Loading JSON from %s...\n", url)
	ethereum, err := parser.LoadJSON(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	if ethereum == nil {
		fmt.Println("Unknown error")
		return
	}

	fmt.Printf("Processing JSON with length of %d\n", len(ethereum.Transactions.Transaction))
	statistics, err := analytics.Process(ethereum)
	if err != nil {
		fmt.Println(err)
		return
	}
	if statistics == nil {
		fmt.Println("Unknown error")
		return
	}

	fmt.Printf("Exporting JSON to %s...\n", output)
	err = parser.ExportJSON(statistics, output)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully!")
}
