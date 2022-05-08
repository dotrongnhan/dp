package adapter

type MacAdapter struct {
	MacMachine Mac
}

func (m *MacAdapter) insertInSquarePort() {
	m.MacMachine.insertInCirclePort()
}
