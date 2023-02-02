// Конкретная реализация.

package main

import (
	"fmt"
	"math/rand"
)

type Email struct {
	Otp
}

func (e *Email) genRandomOTP(length int) string {
	numbers := []rune("123456789")
	randomOTP := make([]rune, length)
	for i := range randomOTP {
		randomOTP[i] += numbers[rand.Intn(len(numbers))]
	}
	fmt.Printf("EMAIL: generating random otp %s\n", string(randomOTP))
	return string(randomOTP)
}

func (e *Email) saveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (e *Email) getMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (e *Email) sendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}
