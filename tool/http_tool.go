package tool

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
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
	LogAPI().Info("GO GET URL: " + req.URL.String())
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

//StructToMap， 将struct对象转换为map
func StructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil
}

//MapToStruct...
func MapToStruct(m map[string]interface{}, strut interface{}) error {
	jsonBody, err := json.Marshal(m)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonBody, &strut); err != nil {
		return err
	}
	return nil
}
