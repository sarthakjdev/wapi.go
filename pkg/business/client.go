package business

import "github.com/sarthakjdev/wapi.go/internal/manager"

type BusinessClient struct {
	PhoneNumber *manager.PhoneNumberManager
	Templates   manager.TemplateManager
}

type BusinessClientConfig struct {
}

func NewBusinessClient(config *BusinessClientConfig) *BusinessClient {
	return &BusinessClient{}
}

func (bc *BusinessClient) GetBusinessId() {
}
