package main

import "visiter/visitor"

func main() {
	square := &visitor.Square{Side: 4}
	circle := &visitor.Circle{Radius: 10}
	rectangle := &visitor.Rectangle{A: 5, B: 4}

	areaCal := &visitor.AreaCalculator{}
	square.Accept(areaCal)
	circle.Accept(areaCal)
	rectangle.Accept(areaCal)
}
