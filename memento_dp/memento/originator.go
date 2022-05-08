package memento

type Originator struct {
	State string
}

func (e *Originator) CreateMemento() *Memento {
	return &Memento{
		State: e.State,
	}
}

func (e *Originator) RestoreMemento(m *Memento) {
	e.State = m.State
}

func (e *Originator) GetState() string {
	return e.State
}

func (e *Originator) SetState(state string) {
	e.State = state
}
