package main

import "template_method/template_method"

func main() {
	smsOTP := &template_method.SMS{}
	o := template_method.OTP{
		ObjOTP: smsOTP,
	}
	o.GenAndSendOTP(1000)

	emailOTP := &template_method.Email{}
	o = template_method.OTP{
		ObjOTP: emailOTP,
	}
	o.GenAndSendOTP(1000)
}
