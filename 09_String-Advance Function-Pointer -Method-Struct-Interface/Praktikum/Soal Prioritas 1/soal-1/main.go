package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) Perkiraan() float64 {
	jarak := c.FuelIn * 1.5
	return jarak
}

func main() {
	myCar := Car{
		Type:   "Sedan",
		FuelIn: 5.0,
	}

	jarak := myCar.Perkiraan()
	fmt.Printf("Jarak perkiraan yang bisa ditempuh oleh mobil %s adalah %.2f mil.\n", myCar.Type, jarak)
}
