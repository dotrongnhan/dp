package command

import "fmt"

type Device interface {
	on()
	off()
}

type TV struct {
	isRunning bool
}

func (t *TV) on() {
	t.isRunning = true
	fmt.Println("Turning TV on")
}

func (t *TV) off() {
	t.isRunning = false
	fmt.Println("Turning TV off")
}