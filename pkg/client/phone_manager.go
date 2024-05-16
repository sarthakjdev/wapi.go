package wapi

type PhoneNumbersManager struct {
	requester requestClient
}

func NewPhoneNumbersManager(requester requestClient) *PhoneNumbersManager {
	return &PhoneNumbersManager{
		requester: requester,
	}
}
