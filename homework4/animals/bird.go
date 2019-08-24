package animals

import "fmt"

// Bird animal
type Bird struct {
}

// Eat function
func (c *Bird) Eat() {
	fmt.Println("Bird eats")
}

// Say function
func (c *Bird) Say() {
	fmt.Println("Bird tweets")
}

// Walk function
func (c *Bird) Walk() {
	fmt.Println("Bird flies")
}
