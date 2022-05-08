package flyweight

import "fmt"

type Soldier struct {
	Name string
}

func NewSoldier(name string) *Soldier {
	return &Soldier{
		Name: name,
	}
}

func (s *Soldier) Promote(context *Context) {
	fmt.Printf("%s %s promoted %d \n", s.Name, context.Id, context.Star)
}
