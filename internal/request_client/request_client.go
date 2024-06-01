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

// RequestClient represents a client for making requests to a cloud API.
type RequestClient struct {
	apiVersion     string
	PhoneNumberId  string
	baseUrl        string
	apiAccessToken string
}

// NewRequestClient creates a new instance of RequestClient.
func NewRequestClient(phoneNumberId string, apiAccessToken string) *RequestClient {
	return &RequestClient{
		apiVersion:     API_VERSION,
		baseUrl:        BASE_URL,
		PhoneNumberId:  phoneNumberId,
		apiAccessToken: apiAccessToken,
	}
}

// RequestCloudApiParams represents the parameters for making a request to the cloud API.
type RequestCloudApiParams struct {
	Body string
	Path string
}

// RequestCloudApi makes a request to the cloud API with the given parameters.
// It returns the response body as a string and any error encountered.
func (requestClientInstance *RequestClient) RequestCloudApi(params RequestCloudApiParams) (string, error) {
	httpRequest, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/%s", REQUEST_PROTOCOL, requestClientInstance.baseUrl, params.Path), strings.NewReader(params.Body))
	if err != nil {
		return "", err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", requestClientInstance.apiAccessToken))
	client := &http.Client{}
	response, err := client.Do(httpRequest)
	if err != nil {
		fmt.Println("Error while requesting cloud api", err)
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("Response from cloud api is", string(body))

	return string(body), nil
}
