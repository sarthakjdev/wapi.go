package request_client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	API_VERSION      = "v20.0"
	BASE_URL         = "graph.facebook.com"
	REQUEST_PROTOCOL = "https"
)

type WhatsappApiType string

const (
	WhatsappApiTypeMessaging WhatsappApiType = "messaging"
	WhatsappApiTypeBusiness  WhatsappApiType = "business"
)

// RequestClient represents a client for making requests to a cloud API.
type RequestClient struct {
	apiVersion     string
	baseUrl        string
	apiAccessToken string
}

// NewRequestClient creates a new instance of RequestClient.
func NewRequestClient(apiAccessToken string) *RequestClient {
	return &RequestClient{
		apiVersion:     API_VERSION,
		baseUrl:        BASE_URL,
		apiAccessToken: apiAccessToken,
	}
}

// RequestCloudApiParams represents the parameters for making a request to the cloud API.
type RequestCloudApiParams struct {
	Body       string
	Path       string
	Method     string
	QueryParam map[string]string
}

// Request makes a request to the cloud API with the given parameters.
// It returns the response body as a string and any error encountered.
func (requestClientInstance *RequestClient) Request(params RequestCloudApiParams) (string, error) {
	queryParamString := ""
	if len(params.QueryParam) > 0 {
		queryParamString = "?"
		for key, value := range params.QueryParam {
			fmt.Println("Key is", key, "Value is", value)
			fmt.Println("queryParamString is", queryParamString)
			// if first query param, don't add "&"
			if queryParamString != "?" {
				queryParamString += "&"
				queryParamString += strings.Join([]string{queryParamString, key, "=", value}, "")
			} else {
				queryParamString += strings.Join([]string{key, "=", value}, "")
			}
		}
	}

	requestPath := strings.Join(
		[]string{REQUEST_PROTOCOL, "://", requestClientInstance.baseUrl, "/", requestClientInstance.apiVersion, "/", params.Path, queryParamString}, "")

	fmt.Println("Requesting cloud api with path", requestPath)

	httpRequest, err := http.NewRequest(params.Method,
		requestPath,
		strings.NewReader(params.Body))
	if err != nil {
		return "", err
	}
	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", requestClientInstance.apiAccessToken))
	httpClient := &http.Client{}
	response, err := httpClient.Do(httpRequest)
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
