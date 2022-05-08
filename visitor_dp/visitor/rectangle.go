package visitor

type Rectangle struct {
	A int
	B int
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}
