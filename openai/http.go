package openai

/**
 * This file contains the http client and request functions for the openai api
 */

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// Constants for HTTP requests

const base_url = "https://api.openai.com"

// add more, e.g. inference, fine tuning new models, etc.

func setRequestHeaders(r *http.Request, apiKey string) *http.Request {
	// Set authentication and content type headers on request
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Bearer "+apiKey)

	return r
}

// Wrapper with authenticated headers
func openaiRequest(method, path string, body io.Reader, key string) (*http.Request, error) {
	req, err := http.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}
	// Add headers to request
	req = setRequestHeaders(req, key)
	return req, nil
}

var httpClient = &http.Client{}

type _OpenAIHTTPClient struct {
	*http.Client
}

// Authenticated HTTP Client for making requests to openai api
var openAIhttp = _OpenAIHTTPClient{httpClient}

// Returns response body for convenience
func (c *_OpenAIHTTPClient) request(method string, url string, body []byte, key string) (*http.Response, []byte, error) {
	r, err := openaiRequest(method, url, bytes.NewBuffer(body), key)

	if err != nil {
		fmt.Println(err)
	}

	// execute request
	resp, err := httpClient.Do(r)

	if err != nil {
		fmt.Println(err)
		// TODO: handle error properly
	}

	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)

	return resp, resBody, err
}
