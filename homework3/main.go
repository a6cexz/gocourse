package main

import (
	"fmt"

	"gocourse/homework3/book"
	_ "gocourse/homework3/cars"
	_ "gocourse/homework3/queue"
)

func main() {
	fmt.Println("Homework 3")

	//cars.Run()
	//queue.Run()
	book.Run()
}
