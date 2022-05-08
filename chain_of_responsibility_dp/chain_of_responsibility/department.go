package chain_of_responsibility

type Department interface {
	Execute(*Patient)
	SetStep(Department)
}
