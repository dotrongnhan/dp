package main

import (
	"flyweight/flyweight"
	"fmt"
	"strconv"
)

func main() {
	factory := flyweight.NewFactory()

	for i := 1; i <= 15; i++ {
		starContext := flyweight.NewContext(strconv.Itoa(i), 5)
		soldier := factory.CreateSoldier("Foot Man")
		soldier.Promote(starContext)
	}

	fmt.Printf("Number of storage soldier map: %d \n", factory.GetSize())
}
