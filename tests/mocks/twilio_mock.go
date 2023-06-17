package mocks


type TwilioMock struct {}

func (twilioMock *TwilioMock) SendSMS(toPhoneNumber string, fromPhoneNumber string, message string) error {
	var error error

	return error
}
