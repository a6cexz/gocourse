package animals

import "fmt"

// Cat animal
type Cat struct {
}

// Eat function
func (c *Cat) Eat() {
	fmt.Println("Cat eats")
}

// Say function
func (c *Cat) Say() {
	fmt.Println("Cat says mayu")
}

// Walk function
func (c *Cat) Walk() {
	fmt.Println("Cat walks")
}
