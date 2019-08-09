package main

import (
	"fmt"
	"math"
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
	fmt.Println("Triangle area, perimeter, hypothenuse")
	a, err := readFloat("Enter a:")
	if err != nil {
		return
	}

	b, err := readFloat("Enter b:")
	if err != nil {
		return
	}

	area := a * b / 2.0
	c := math.Sqrt(a*a + b*b)
	perimeter := a + b + c

	fmt.Printf("Triangle sides a: %.2f b: %.2f\n", a, b)
	fmt.Printf("area: %.2f\n", area)
	fmt.Printf("perimeter: %.2f\n", perimeter)
	fmt.Printf("hypothenuse: %.2f\n", c)
}
