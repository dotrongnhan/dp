package abstractFactory

type InterfaceSportsFactory interface {
	MakeShoe() InterfaceShoe
	MakeShort() InterfaceShort
}

func GetSportFactory(brand string) InterfaceSportsFactory {
	switch brand {
	case "adidas":
		return &Adidas{}
	default:
		return nil
	}
}
