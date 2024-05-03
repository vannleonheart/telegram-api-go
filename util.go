package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func sendHttpRequest(method, url string, data interface{}, headers *map[string]string, result interface{}) error {
	var requestBody io.Reader
	switch method {
	case http.MethodGet:
		if data != nil {
			q := generateQueryString(data)
			url = fmt.Sprintf("%s?%s", url, q)
		}
	case http.MethodPost:
		if data != nil {
			reqBody, err := generateRequestBody(data, headers)
			if err != nil {
				return err
			}

			requestBody = *reqBody
		}
	}

	cl := &http.Client{}

	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return err
	}

	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}

	resp, err := cl.Do(req)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	byteBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteBody, &result)
}

func sendHttpGet(url string, params interface{}, headers *map[string]string, result interface{}) error {
	return sendHttpRequest(http.MethodGet, url, params, headers, result)
}

func sendHttpPost(url string, data interface{}, headers *map[string]string, result interface{}) error {
	return sendHttpRequest(http.MethodPost, url, data, headers, result)
}

func generateRequestBody(data interface{}, headers *map[string]string) (*io.Reader, error) {
	contentType := ""
	if headers != nil {
		cType, exist := (*headers)["Content-Type"]
		if exist {
			contentType = cType
		}
	}

	var result io.Reader

	switch contentType {
	default:
		bytePostData, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		result = bytes.NewBuffer(bytePostData)
	case "application/x-www-form-urlencoded":
		encodedData := generateQueryString(data)
		result = strings.NewReader(encodedData)
	}

	return &result, nil
}

func generateQueryString(data interface{}) string {
	by, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	var jsonData map[string]interface{}

	if err = json.Unmarshal(by, &jsonData); err != nil {
		return ""
	}

	values := url.Values{}
	for k, v := range jsonData {
		values.Add(k, fmt.Sprintf("%+v", v))
	}

	return values.Encode()
}
