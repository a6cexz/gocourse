package main

import (
	"fmt"
	"strconv"
)

func readFloat(prompt string) (float64, error) {
	fmt.Println(prompt)

	var str string
	_, err := fmt.Scanln(&str)
	if err != nil {
		fmt.Printf("Error reading value: %v", err)
		return -1, err
	}

	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Printf("Error parsing float value: %v", err)
		return -1, err
	}

	return val, nil
}

func main() {
	fmt.Println("Bank")
	amount, err := readFloat("Enter amount:")
	if err != nil {
		return
	}

	percent, err := readFloat("Enter year percent:")
	if err != nil {
		return
	}

	rate := percent / 100.0

	year1 := amount + amount*rate
	year2 := year1 + year1*rate
	year3 := year2 + year2*rate
	year4 := year3 + year3*rate
	year5 := year4 + year4*rate

	fmt.Printf("Start amount: %.2f percent: %.2f\n", amount, percent)
	fmt.Printf("Year 5 amount: %.2f\n", year5)

	fmt.Println()
	fmt.Println("Details by year:")
	fmt.Printf("Year 1 amount: %.2f\n", year1)
	fmt.Printf("Year 2 amount: %.2f\n", year2)
	fmt.Printf("Year 3 amount: %.2f\n", year3)
	fmt.Printf("Year 4 amount: %.2f\n", year4)
	fmt.Printf("Year 5 amount: %.2f\n", year5)
}
