package animals

import "fmt"

// Dog animal
type Dog struct {
}

// Eat function
func (c *Dog) Eat() {
	fmt.Println("Dog eats")
}

// Say function
func (c *Dog) Say() {
	fmt.Println("Dog says woof")
}

// Walk function
func (c *Dog) Walk() {
	fmt.Println("Dog runs")
}
