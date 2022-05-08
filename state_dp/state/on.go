package state

import "fmt"

type on struct {
}

func newON() state {
	return &on{}
}

func (o *on) off(m *Machine) {
	fmt.Println("Machine turned off")
	m.setCurrent(newOFF())
}

func (o *on) on(m *Machine) {
	fmt.Println("Machine is turning on")
}
