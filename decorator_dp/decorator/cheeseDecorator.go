package decorator

type CheeseDecorator struct {
	Pizza InterfacePizza
}

func NewCheeseDecorator(pizza InterfacePizza) *CheeseDecorator {
	return &CheeseDecorator{
		pizza,
	}
}

func (c *CheeseDecorator) DoPizza() string {
	pizzaType := c.Pizza.DoPizza()
	return pizzaType + addFlavour()
}

func addFlavour() string {
	return "Cheese"
}
