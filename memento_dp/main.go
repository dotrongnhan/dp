package main

import (
	"fmt"
	"mementor/memento"
)

func main() {
	careTaker := &memento.CareTaker{
		MementoArray: make([]*memento.Memento, 0),
	}

	originator := &memento.Originator{
		State: "A",
	}
	fmt.Printf("Originator current state: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator current state: %s\n", originator.GetState())

	careTaker.AddMemento(originator.CreateMemento())
	originator.SetState("C")

	fmt.Printf("Originator current state: %s\n", originator.GetState())
	careTaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(careTaker.GetMemento(1))
	fmt.Printf("Originator current state: %s\n", originator.GetState())
}
