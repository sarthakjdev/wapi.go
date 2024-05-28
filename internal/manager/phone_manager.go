package manager

type PhoneNumbersManager struct {
	requester RequestClient
}

func NewPhoneNumbersManager(requester RequestClient) *PhoneNumbersManager {
	return &PhoneNumbersManager{
		requester: requester,
	}
}
