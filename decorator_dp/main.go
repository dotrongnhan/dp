package main

import (
	"decorator/decorator"
	"fmt"
)

func main() {
	tomato := &decorator.TomatoPizza{}
	chicken := &decorator.ChickenPizza{}

	fmt.Println(tomato.DoPizza())
	fmt.Println(chicken.DoPizza())

	cheeseDecorator := decorator.NewCheeseDecorator(chicken)
	fmt.Println(cheeseDecorator.DoPizza())
}
