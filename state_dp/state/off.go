package state

import "fmt"

type off struct {
}

func newOFF() state {
	return &off{}
}

func (o *off) on(m *Machine) {
	fmt.Println("Machine turned on")
	m.setCurrent(newON())
}

func (o *off) off(m *Machine) {
	fmt.Println("Machine is turning off")
}
