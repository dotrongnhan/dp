package chain_of_responsibility

import "fmt"

type Cashier struct {
	next Department
}

func (r *Cashier) Execute(p *Patient) {
	if p.isPaid {
		fmt.Println("Patient has paid his bill already")
		r.next.Execute(p)
		return
	}
	fmt.Println("Patient is paying his bill")
	p.isPaid = true
}

func (r *Cashier) SetStep(next Department) {
	r.next = next
}
