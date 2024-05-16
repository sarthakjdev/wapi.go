package manager

import (
	"fmt"
)

type requestClient struct {
}

func NewRequestClient() *requestClient {
	return &requestClient{}
}

func (requestClientInstance *requestClient) requestCloudApi() {
	fmt.Println("Requesting cloud api")
}
