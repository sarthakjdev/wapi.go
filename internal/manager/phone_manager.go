package manager

type PhoneNumbersManager struct {
	requester requestClient
}

func NewPhoneNumbersManager(requester requestClient) *PhoneNumbersManager {
	return &PhoneNumbersManager{
		requester: requester,
	}
}
