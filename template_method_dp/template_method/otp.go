package template_method

type OTP struct {
	ObjOTP InterfaceOTP
}

func (o *OTP) GenAndSendOTP(otpLength int) error {
	otp := o.ObjOTP.GenRandomOTP(otpLength)
	o.ObjOTP.SaveOTPCache(otp)
	message := o.ObjOTP.GetMessage(otp)
	err := o.ObjOTP.SendNotification(message)
	if err != nil {
		return err
	}
	return nil
}
