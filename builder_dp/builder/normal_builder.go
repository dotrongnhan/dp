package builder

type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (n *normalBuilder) setWindowType() {
	n.windowType = "Wooden"
}

func (n *normalBuilder) setDoorType() {
	n.doorType = "Wooden"
}

func (n *normalBuilder) setNumFloor() {
	n.floor = 4
}

func (n *normalBuilder) getHouse() House {
	return House{
		doorType:   n.doorType,
		windowType: n.windowType,
		floor:      n.floor,
	}
}
