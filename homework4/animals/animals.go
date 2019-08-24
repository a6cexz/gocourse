package animals

import (
	"fmt"
)

// Animal interface
type Animal interface {
	Eat()
	Say()
	Walk()
}

// Run functions shows animals demo
func Run() {
	fmt.Println()
	fmt.Println("Animals Demo")
	fmt.Println()

	var animals = []Animal{
		&Cat{},
		&Dog{},
		&Bird{},
	}

	for _, a := range animals {
		a.Eat()
		a.Say()
		a.Walk()
	}
}
