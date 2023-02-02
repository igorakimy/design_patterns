// Конкретная реализация.

package main

import (
	"fmt"
	"math/rand"
)

type Sms struct {
	Otp
}

func (s *Sms) genRandomOTP(length int) string {
	numbers := []rune("123456789")
	randomOTP := make([]rune, length)
	for i := range randomOTP {
		randomOTP[i] += numbers[rand.Intn(len(numbers))]
	}
	fmt.Printf("SMS: generating random otp %s\n", string(randomOTP))
	return string(randomOTP)
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
