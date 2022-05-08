package visitor

import "fmt"

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) VisitForSquare(s *Square) {
	fmt.Println("Calculating area for Square")
}

func (a *AreaCalculator) VisitForCircle(c *Circle) {
	fmt.Println("Calculating area for Circle")
}

func (a *AreaCalculator) VisitForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for Rectangle")
}
