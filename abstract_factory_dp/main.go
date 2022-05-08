package main

import (
	"abstract_factory/abstractFactory"
	"fmt"
)

func main() {
	adidasFactory := abstractFactory.GetSportFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	adidasShort := adidasFactory.MakeShort()
	PrintValue(adidasShoe, adidasShort)
}

func PrintValue(shoe abstractFactory.InterfaceShoe, short abstractFactory.InterfaceShort) {
	fmt.Printf("Size Adidas Shoe is %v\n", shoe.GetSize())
	fmt.Printf("Logo Adidas Shoe is %v\n", shoe.GetLogo())
	fmt.Printf("Size Adidas Short is %v\n", short.GetSize())
	fmt.Printf("Logo Adidas Short is %v", short.GetLogo())
}
