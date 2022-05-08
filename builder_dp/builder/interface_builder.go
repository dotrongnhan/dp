package builder

type InterfaceBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func GetBuilder(builderType string) InterfaceBuilder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	default:
		return nil
	}
}
