package config

import "os"

type Env struct {
	TwilioAccountSid string
	TwilioAuthToken string
	TwilioPhoneNumber string
	TwimlUrl string
}

func GetEnv() *Env {
	return &Env{
		TwilioAccountSid: os.Getenv("TWILIO_ACCOUNT_SID"),
		TwilioAuthToken: os.Getenv("TWILIO_AUTH_TOKEN"),
		TwilioPhoneNumber: os.Getenv("TWILIO_PHONE_NUMBER"),
		TwimlUrl: os.Getenv("TWIML_URL"),
	}
}