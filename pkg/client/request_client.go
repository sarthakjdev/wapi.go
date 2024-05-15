package client

import "fmt"

type requestClient struct {
}

func NewRequestClient() *requestClient {
	return &requestClient{}
}

func (*requestClient) requestCloudApi() {
	fmt.Println("Requesting cloud api")
}
