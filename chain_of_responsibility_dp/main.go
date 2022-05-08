package main

import "chain_of_responsibility/chain_of_responsibility"

func main() {
	cashier := &chain_of_responsibility.Cashier{}

	medical := &chain_of_responsibility.Medical{}
	medical.SetStep(cashier)

	doctor := &chain_of_responsibility.Doctor{}
	doctor.SetStep(medical)

	reception := &chain_of_responsibility.Reception{}
	reception.SetStep(doctor)

	patient := &chain_of_responsibility.Patient{Name: "Beo"}
	reception.Execute(patient)
}
