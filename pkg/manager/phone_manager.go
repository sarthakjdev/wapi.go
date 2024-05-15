package manager

type PhoneNumbersManager struct {
	BaseManager
}

func (*PhoneNumbersManager) NewPhoneNumbersManager() *PhoneNumbersManager {
	return &PhoneNumbersManager{}
}
