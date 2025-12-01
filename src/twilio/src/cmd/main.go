package main

import (
	"sample.twilio/config"
	"sample.twilio/service"
)

func main() {
	env := config.GetEnv()

	// change your phone number
	phoneNumber := "+821012345678"
	service.Calling(env, phoneNumber)
}