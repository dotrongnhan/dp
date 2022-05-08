package main

import "state/state"

func main() {
	machine := state.NewMachine()
	machine.ON()
	machine.ON()
	machine.OFF()
	machine.OFF()

}
