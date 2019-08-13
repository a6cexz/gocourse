package main

import (
	"fmt"
)

func division() {
	for i := 1; i <= 10; i++ {
		if isEven(i) {
			fmt.Printf("%d is even\n", i)
		}

		if isDivisibleBy3(i) {
			fmt.Printf("%d is divisible by 3\n", i)
		}
	}
}

func isEven(n int) bool {
	return n%2 == 0
}

func isDivisibleBy3(n int) bool {
	return n%3 == 0
}

func fib() {
	var fn1 int64
	var fn int64
	for i := 0; i < 100; i++ {
		if i <= 1 {
			fn1 = 0
			fn = 1
		}
		f := fn1 + fn
		fmt.Printf("%d ", f)
		fn1 = fn
		fn = f
	}
}

func primes(n int, max int) {
	count := 0
	primes := []int{}
	crossed := make([]bool, n)
	for i := 2; i < n; i++ {
		if crossed[i] {
			continue
		}
		primes = append(primes, i)
		count++
		if count >= max {
			break
		}
		for j := i + i; j < n; j += i {
			crossed[j] = true
		}
	}

	fmt.Printf("list of first %d primes:\n", len(primes))
	for i := 0; i < len(primes); i++ {
		fmt.Printf("%d ", primes[i])
	}
}

func main() {
	fmt.Println("Homework 2")

	fmt.Println()
	fmt.Println("Division")
	division()

	fmt.Println()
	fmt.Println("Fib")
	fib()

	fmt.Println()
	fmt.Println("Primes")
	primes(1000, 100)
}
