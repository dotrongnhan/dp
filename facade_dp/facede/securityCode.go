package facede

import "fmt"

type SecurityCode struct {
	Code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		Code: code,
	}
}

func (s *SecurityCode) CheckCode(code int) error {
	if s.Code != code {
		return fmt.Errorf("Security code is incorrect")
	}
	fmt.Println("Security code verified")
	return nil
}
