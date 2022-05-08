package chain_of_responsibility

type Patient struct {
	Name               string
	isRegistered       bool
	isDoctorChecked    bool
	isMedicineProvided bool
	isPaid             bool
}
