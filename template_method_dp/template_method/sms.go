package template_method

import (
	"fmt"
	"math/rand"
	"strconv"
)

type SMS struct {
}

func (s *SMS) GenRandomOTP(len int) string {
	randomOTP := rand.Intn(10000)
	fmt.Printf("SMS: generating random OTP %v \n", randomOTP)
	return strconv.Itoa(randomOTP)
}

func (s *SMS) SaveOTPCache(otp string) {
	fmt.Printf("SMS: saving OTP %v to cache \n", otp)
}

func (s *SMS) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *SMS) SendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %v\n", message)
	return nil
}
