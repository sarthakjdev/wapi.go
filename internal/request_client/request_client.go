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

func (client *RequestClient) NewBusinessApiRequest(path, method string) *WhatsappBusinessManagementApiRequest {
	return &WhatsappBusinessManagementApiRequest{
		Path:        path,
		Fields:      []BusinessApiRequestQueryParamField{},
		Requester:   client,
		Method:      method,
		QueryParams: map[string]string{},
	}
}

type BusinessApiRequestQueryParamField struct {
	Name    string
	Filters map[string]string
}

func (field *BusinessApiRequestQueryParamField) AddFilter(key, value string) {
	field.Filters[key] = value
}

type WhatsappBusinessManagementApiRequest struct {
	Path        string
	Method      string
	Body        string
	Fields      []BusinessApiRequestQueryParamField
	QueryParams map[string]string
	Requester   *RequestClient
}

func (request *WhatsappBusinessManagementApiRequest) AddField(field BusinessApiRequestQueryParamField) *BusinessApiRequestQueryParamField {
	// * NOTE:  when we say we need to add a field to the request, it means we need to add a query param to the request
	// * note that if there need to be multiple fields in a single request then the list of fields should be command separated
	// * for example: fields=field1,field2,field3
	// * also note that if there filters in a field then they should be called like a function in the param string, for ex: fields=field1.filter1(value1).filter2(value2),field2.filter1(value1)
	request.Fields = append(request.Fields, field)
	return &field
}

func (request *WhatsappBusinessManagementApiRequest) AddQueryParam(key, value string) {
	request.QueryParams[key] = value
}

func (request *WhatsappBusinessManagementApiRequest) SetMethod(method string) {
	request.Method = method
}

func (request *WhatsappBusinessManagementApiRequest) SetBody(body string) {
	request.Body = body
}

func (businessRequest *WhatsappBusinessManagementApiRequest) Execute() (string, error) {
	// check if there are any fields in the request
	var queryParam = map[string]string{}
	if len(businessRequest.Fields) > 0 {
		fieldsString := ""
		for _, field := range businessRequest.Fields {
			newFieldString := ""
			if fieldsString != "" {
				newFieldString = ","
			}
			filterString := ""
			for key, value := range field.Filters {
				filterString += strings.Join([]string{".", key, "(", value, ")"}, "")
			}
			newFieldString = strings.Join([]string{field.Name, filterString}, "")
			fieldsString += newFieldString
		}

		queryParam["fields"] = fieldsString
	}

	if len(businessRequest.QueryParams) > 0 {
		for key, value := range businessRequest.QueryParams {
			queryParam[key] = value
		}
	}

	fmt.Println("Query params are", queryParam)

	response, err := businessRequest.Requester.Request(RequestCloudApiParams{
		Path:       businessRequest.Path,
		Body:       businessRequest.Body,
		Method:     businessRequest.Method,
		QueryParam: queryParam,
	})

	if err != nil {
		fmt.Println("Error while executing business api request", err)
		return "", nil
	}

	fmt.Println("Response from business api is", response)
	return response, err
}
