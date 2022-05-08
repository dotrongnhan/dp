package abstractFactory

type Adidas struct {
}

func (a *Adidas) MakeShoe() InterfaceShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "Adidas",
			size: 13,
		},
	}
}

func (a *Adidas) MakeShort() InterfaceShort {
	return &adidasShort{
		short: short{
			logo: "Adidas",
			size: 10,
		},
	}
}
