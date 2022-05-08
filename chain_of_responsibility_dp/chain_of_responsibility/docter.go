package chain_of_responsibility

import "fmt"

type Doctor struct {
	next Department
}

func (r *Doctor) Execute(p *Patient) {
	if p.isDoctorChecked {
		fmt.Println("Patient has checked by Doctor already")
		r.next.Execute(p)
		return
	}
	fmt.Println("Doctor is checking patient")
	p.isDoctorChecked = true
	r.next.Execute(p)
}

func (r *Doctor) SetStep(next Department) {
	r.next = next
}
