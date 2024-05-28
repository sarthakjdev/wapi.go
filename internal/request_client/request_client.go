package requestclient

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	API_VERSION      = "v19.0"
	BASE_URL         = "graph.facebook.com"
	REQUEST_PROTOCOL = "https"
)

type RequestClient struct {
	apiVersion     string
	PhoneNumberId  string
	baseUrl        string
	apiAccessToken string
}

func NewRequestClient(phoneNumberId string, apiAccessToken string) *RequestClient {
	return &RequestClient{
		apiVersion:     API_VERSION,
		baseUrl:        BASE_URL,
		PhoneNumberId:  phoneNumberId,
		apiAccessToken: apiAccessToken,
	}
}

type RequestCloudApiParams struct {
	Body string
	Path string
}

func (requestClientInstance *RequestClient) RequestCloudApi(params RequestCloudApiParams) {
	httpRequest, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/%s", REQUEST_PROTOCOL, requestClientInstance.baseUrl, params.Path), strings.NewReader(params.Body))
	if err != nil {
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", requestClientInstance.apiAccessToken))
	client := &http.Client{}
	response, err := client.Do(httpRequest)
	if err != nil {
		fmt.Println("Error while requesting cloud api", err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println("Response from cloud api is", string(body))
}
