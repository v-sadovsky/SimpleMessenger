package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Do(method string, endpoint string, params map[string]string, data any) (int, []byte, error) {
	handle := func(err error) (int, []byte, error) {
		return http.StatusInternalServerError, nil, fmt.Errorf("doing %s request to the %s endpoint: %w", method, endpoint, err)
	}

	var err error
	var payload []byte
	queryString := endpoint

	switch method {
	case http.MethodGet, http.MethodDelete:
		payload = nil
	case http.MethodPost, http.MethodPut:
		payload, err = json.Marshal(data)
		if err != nil {
			return handle(err)
		}
	default:
		return handle(fmt.Errorf("method %s unsupported", method))
	}

	if params != nil {
		query := url.Values{}
		for key, value := range params {
			query.Add(key, value)
		}
		queryString = fmt.Sprintf("%s?%s", endpoint, query.Encode())
	}

	var body io.Reader
	if payload != nil {
		body = bytes.NewReader(payload)
	}

	request, err := http.NewRequest(method, queryString, body)
	if err != nil {
		return handle(err)
	}

	//request.SetBasicAuth(bm.Email, bm.Password)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return handle(err)
	}
	defer func() { _ = resp.Body.Close() }()

	//if resp.StatusCode != http.StatusOK {
	//	var msg string
	//	respBody, err := io.ReadAll(resp.Body)
	//	if err != nil {
	//		return handle(fmt.Errorf("%w: response status code is `%s`", err, resp.Status))
	//	}
	//
	//	msg = fmt.Sprintf(": %s", string(respBody))
	//	return handle(fmt.Errorf("response status code is `%s`%s", resp.Status, msg))
	//}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return handle(err)
	}

	return resp.StatusCode, respBody, nil
}
