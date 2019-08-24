package main

import (
	"fmt"

	"gocourse/homework4/animals"
	"gocourse/homework4/calculator"
	"gocourse/homework4/phones"
)

func main() {
	fmt.Println("Homework 4")

	animals.Run()
	phones.Run()
	calculator.Run()
}
