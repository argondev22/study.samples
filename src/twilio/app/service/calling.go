package service

import (
	"log"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"

	"sample.twilio/config"
)

func Calling(env *config.Env, phoneNumber string) {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: env.TwilioAccountSid,
		Password: env.TwilioAuthToken,
	})

	params := &twilioApi.CreateCallParams{
		Url:  &env.TwimlUrl,
		From: &env.TwilioPhoneNumber,
		To:   &phoneNumber,
	}

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		log.Println("Failed to create call", err)
		return
	}

	log.Printf("Make a Call: From %s to %s", env.TwilioPhoneNumber, phoneNumber)
	log.Printf("Call Created: %+v", resp)
}