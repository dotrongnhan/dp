package main

import "observer/observer"

func main() {
	shirtItem := observer.NewItem("shirt")
	customer1 := &observer.Customer{ID: "AA1"}
	customer2 := &observer.Customer{ID: "AB1"}
	shirtItem.Register(customer1)
	shirtItem.Register(customer2)
	shirtItem.DeRegister(customer1)
	shirtItem.UpdateAvailability()
}
