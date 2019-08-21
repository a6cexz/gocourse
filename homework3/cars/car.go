package cars

import (
	"fmt"
)

// Car struct stores car data
type Car struct {
	Brand           string
	Year            int
	IsTruck         bool
	CargoCapacity   int
	CargoSize       int
	IsEngineStarted bool
	AreWindowsOpen  bool
}

func printCar(car *Car) {
	fmt.Println("Car Info")
	fmt.Printf("Brand: %v Year: %v IsTruck: %v CargoCapacity: %v CargoSize: %v IsEngineStarted: %v AreWindowsOpen:%v\n",
		car.Brand,
		car.Year,
		car.IsTruck,
		car.CargoCapacity,
		car.CargoSize,
		car.IsEngineStarted,
		car.AreWindowsOpen)
	fmt.Println()
}

// Run shows cars demo
func Run() {
	fmt.Println()
	fmt.Println("Cars")

	ford := &Car{
		Brand:           "Ford",
		Year:            2001,
		IsTruck:         false,
		CargoCapacity:   200,
		CargoSize:       100,
		IsEngineStarted: true,
		AreWindowsOpen:  false,
	}
	printCar(ford)

	caterpillar := &Car{
		Brand:           "Caterpillar797",
		Year:            1998,
		IsTruck:         true,
		CargoCapacity:   32700,
		CargoSize:       1000,
		IsEngineStarted: false,
		AreWindowsOpen:  true,
	}
	printCar(caterpillar)
}
