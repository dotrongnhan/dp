package main

import (
	"builder/builder"
	"fmt"
)

func main() {
	normalBuilder := builder.GetBuilder("normal")
	director := builder.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()
	fmt.Println(normalHouse)
}
