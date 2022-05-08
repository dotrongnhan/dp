package main

import "mediator/mediator"

func main()  {
	stationManager := mediator.NewStationManager()
	passengerTrain := &mediator.PassengerTrain{
		Mediator: stationManager,
	}
	goodsTrain := &mediator.GoodsTrain{
		Mediator: stationManager,
	}

	passengerTrain.RequestArrival()
	goodsTrain.RequestArrival()
	passengerTrain.Departure()

}