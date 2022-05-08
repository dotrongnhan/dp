package mediator

import "fmt"

type PassengerTrain struct {
	Mediator	Mediator
}

func (p *PassengerTrain) Departure() {
	fmt.Println("Passenger train: Leaving")
	p.Mediator.NotifyFree()
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Println("Passenger train: Arrival permitted. Landing")
}

func (p *PassengerTrain) RequestArrival() {
	if p.Mediator.CanLand(p) {
		fmt.Println("Passenger train: Landing")
		return
	}
	fmt.Println("Passenger train: Waiting")
}