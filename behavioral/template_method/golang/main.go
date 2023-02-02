// Клиентский код.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	smsOTP := &Sms{}
	o := Otp{
		IOtp: smsOTP,
	}
	o.genAndSendOTP(4)

	fmt.Println()
	emailOTP := &Email{}
	o = Otp{
		IOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}

// SMS: generating random otp 3266
// SMS: saving otp: 3266 to cache
// SMS: sending sms: SMS OTP for login is 3266
//
// EMAIL: generating random otp 4154
// EMAIL: saving otp: 4154 to cache
// EMAIL: sending email: EMAIL OTP for login is 4154
