package chain_of_responsibility

import "fmt"

type Medical struct {
	next Department
}

func (r *Medical) Execute(p *Patient) {
	if p.isMedicineProvided {
		fmt.Println("Patient has provided medicine already")
		r.next.Execute(p)
		return
	}
	fmt.Println("We are providing medicine to patient")
	p.isMedicineProvided = true
	r.next.Execute(p)
}

func (r *Medical) SetStep(next Department) {
	r.next = next
}
