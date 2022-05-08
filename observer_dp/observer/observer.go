package observer

type Observer interface {
	Update(str string)
	GetId() string
}
