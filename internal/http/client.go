package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func formatURL(url string) string {
	if url[len(url)-1] != '?' {
		url += "?"
	}
	return url
}

func GetAndDecode(
	url string,
	headers map[string]string,
	params map[string]string,
	response interface{},
) error {
	url = formatURL(url)

	for key, value := range params {
		url += key + "=" + value + "&"
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	httpResponse, err := client.Do(req)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK &&
		httpResponse.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", httpResponse.StatusCode)
	}

	bytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, response)
}

func PostAndDecode(
	url string,
	headers map[string]string,
	body interface{},
	response interface{},
) error {
	url = formatURL(url)

	client := &http.Client{}

	reqBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	httpResponse, err := client.Do(req)
	if err != nil {
		return err
	}

	defer httpResponse.Body.Close()

	if httpResponse.StatusCode != http.StatusOK &&
		httpResponse.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", httpResponse.StatusCode)
	}

	bytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, response)
}
