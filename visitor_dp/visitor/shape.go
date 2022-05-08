package visitor

type Shape interface {
	accept(Visitor)
}
