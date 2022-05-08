package visitor

type Visitor interface {
	VisitForSquare(s *Square)
	VisitForCircle(c *Circle)
	VisitForRectangle(r *Rectangle)
}
