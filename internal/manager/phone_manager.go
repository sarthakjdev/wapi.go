package manager

import requestclient "github.com/sarthakjdev/wapi.go/internal/request_client"

type PhoneNumbersManager struct {
	requester requestclient.RequestClient
}

func NewPhoneNumbersManager(requester requestclient.RequestClient) *PhoneNumbersManager {
	return &PhoneNumbersManager{
		requester: requester,
	}
}
