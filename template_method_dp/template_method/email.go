package template_method

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Email struct {
}

func (s *Email) GenRandomOTP(len int) string {
	randomOTP := rand.Intn(len)
	fmt.Printf("Email: generating random OTP %v \n", randomOTP)
	return strconv.Itoa(randomOTP)
}

func (s *Email) SaveOTPCache(otp string) {
	fmt.Printf("Email: saving OTP %v to cache \n", otp)
}

func (s *Email) GetMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (s *Email) SendNotification(message string) error {
	fmt.Printf("Email: sending mail:  %v\n", message)
	return nil
}
