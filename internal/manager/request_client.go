package manager

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	API_VERSION      = "v.19"
	BASE_URL         = "graph.facebook.com"
	REQUEST_PROTOCOL = "https"
)

type RequestClient struct {
	apiVersion     string
	phoneNumberId  string
	baseUrl        string
	apiAccessToken string
}

func NewRequestClient(phoneNumberId string, apiAccessToken string) *RequestClient {
	return &RequestClient{
		apiVersion:     API_VERSION,
		baseUrl:        BASE_URL,
		phoneNumberId:  phoneNumberId,
		apiAccessToken: apiAccessToken,
	}
}

type requestCloudApiParams struct {
	body string
	path string
}

func (requestClientInstance *RequestClient) requestCloudApi(params requestCloudApiParams) {
	httpRequest, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/%s", REQUEST_PROTOCOL, requestClientInstance.baseUrl, params.path), strings.NewReader(params.body))
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
