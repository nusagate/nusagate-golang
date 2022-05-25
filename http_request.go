package nusagate

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpRequest interface {
	Call(method string, url string, secretKey string, body io.Reader, result interface{}) *Error
}

type HttpRequestImpl struct {
	Config *ConfigOption
}

func (client *HttpRequestImpl) Call(method string, url string, apiKey string, secretKey string, body io.Reader, result interface{}) *Error {
	if apiKey == "" {
		return &Error{
			Status:  http.StatusUnauthorized,
			Message: "missing/invalid api key",
		}
	}

	if secretKey == "" {
		return &Error{
			Status:  http.StatusUnauthorized,
			Message: "missing/invalid secret key",
		}
	}

	req, errNewReq := http.NewRequest(method, url, body)
	if errNewReq != nil {
		return &Error{
			Status:   http.StatusInternalServerError,
			Message:  "Failed create request",
			RawError: errNewReq.Error(),
		}
	}

	encodedAuth := base64.StdEncoding.EncodeToString([]byte(apiKey + ":" + secretKey))
	req.Header.Add("Authorization", "Basic "+encodedAuth)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return client.doRequest(req, result)
}

func (client *HttpRequestImpl) doRequest(req *http.Request, result interface{}) *Error {
	httpClient := &http.Client{}

	response, errRequest := httpClient.Do(req)
	if errRequest != nil {
		return ErrorGo(errRequest)
	}
	defer response.Body.Close()

	respBody, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		return ErrorGo(errRead)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return ErrorHttp(response.StatusCode, respBody)
	}

	errUnmarshall := json.Unmarshal(respBody, &result)
	if errUnmarshall != nil {
		return ErrorGo(errUnmarshall)
	}

	return nil
}
