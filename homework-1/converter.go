package main

import (
	"fmt"
	"strconv"
)

const rubToUsd = 65.09

func main() {
	fmt.Println("RUB to USD converter")
	fmt.Println("Enter RUB value:")

	var rubStr string
	_, err := fmt.Scanln(&rubStr)
	if err != nil {
		fmt.Printf("Error reading value: %v", err)
		return
	}

	rub, err := strconv.ParseFloat(rubStr, 64)
	if err != nil {
		fmt.Printf("Error parsing float value: %v", err)
		return
	}

	usd := rub / rubToUsd
	fmt.Printf("%.2f RUB converted to USD: $%.2f", rub, usd)
}
