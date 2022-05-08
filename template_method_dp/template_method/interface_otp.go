package template_method

type InterfaceOTP interface {
	GenRandomOTP(int) string
	SaveOTPCache(otp string)
	GetMessage(str string) string
	SendNotification(message string) error
}
