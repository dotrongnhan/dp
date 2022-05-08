package builder

type Director struct {
	builder InterfaceBuilder
}

func NewDirector(b InterfaceBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) BuildHouse() House {
	d.builder.setNumFloor()
	d.builder.setWindowType()
	d.builder.setDoorType()
	return d.builder.getHouse()
}
