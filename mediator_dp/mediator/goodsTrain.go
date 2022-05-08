package mediator

import "fmt"

type GoodsTrain struct {
	Mediator	Mediator
}

func (p *GoodsTrain) Departure() {
	fmt.Println("Passenger train: Leaving")
	p.Mediator.NotifyFree()
}

func (p *GoodsTrain) PermitArrival() {
	fmt.Println("Passenger train: Arrival permitted. Landing")
}

func (p *GoodsTrain) RequestArrival() {
	if p.Mediator.CanLand(p) {
		fmt.Println("Passenger train: Landing")
		return
	}
	fmt.Println("Passenger train: Waiting")
}

