package tool

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// HTTPGet...
func HTTPGet(url string, params map[string]interface{}, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		LogAPI().Error(err)
		return nil, err
	}
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val.(string))
		}
		req.URL.RawQuery = q.Encode()
	}

	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}

	client := &http.Client{}
	//LogAPI().Info("GO GET URL: " + req.URL.String())
	return client.Do(req)
}

// HTTPPost...
func HTTPPost(url string, body map[string]interface{}, params map[string]interface{},
	headers map[string]interface{}) (*http.Response, error) {

	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			LogAPI().Error(err)
			return nil, err
		}
	}
	LogAPI().Info(string(bodyJson))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		LogAPI().Error(err)
		return nil, err
	}
	req.Header.Set("Content-type", "application/json")

	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val.(string))
		}
		req.URL.RawQuery = q.Encode()
	}
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val.(string))
		}
	}

	client := &http.Client{}
	LogAPI().Info("GO POST URL: " + req.URL.String())
	return client.Do(req)
}
