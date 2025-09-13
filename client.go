// Package client provides a specialized HTTP client for interacting with the Data Warehouse microservice.
// This package handles all aspects of communication with the microservice, ensuring proper data exchange and
// streamlining the process of sending and receiving data. The client supports both GET and POST HTTP methods,
// and is specifically tailored to handle JSON formatted data, aligning with the microservice's API specifications.
//
// The Client struct central to this package is configured to authenticate via an API key and connect to the
// microservice using a provided base URL. It encapsulates functionality for constructing requests, setting
// appropriate headers, and parsing responses, thus abstracting the complexity of direct HTTP interactions.
//
// Primary functionalities include:
//   - Initialization of a new client instance with configuration for URL and API key.
//   - Execution of GET requests to retrieve data from the microservice using well-defined endpoints.
//   - Execution of POST requests to submit data to the microservice, facilitating real-time data processing and updates.
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Client is a struct that encapsulates necessary details and methods to interact with the Data Warehouse microservice.
// This struct is designed to manage the HTTP communication with the microservice, handling tasks such as building requests,
// sending them, and processing the responses. Client makes it easy to integrate with the microservice without worrying
// about the underlying HTTP operations.
//
// Fields:
//   - url: Base URL of the Data Warehouse microservice. This is the root address to which all API endpoints are appended.
//   - client: A pointer to an http.Client that performs the actual HTTP requests. This allows for customization of aspects
//     like timeouts and redirection policies.
//   - apiKey: API key used for authentication with the microservice. This key is essential to ensure secure access to the API
//     and is sent as a header in each request to authenticate the client.
//
// The design of the Client struct emphasizes ease of use and flexibility, enabling developers to interact with the microservice
// efficiently while maintaining high standards of security.
type Client struct {
	url    string
	client *http.Client
	apiKey string
}

// New initializes and returns a new Client instance.
// This function will return an error if the URL or API key are not provided.
//
// Parameters:
//   - url: Base URL of the API
//   - apiKey: API key for authenticating requests
//
// Returns:
//   - *Client: A pointer to the newly created Client instance
//   - error: Error if the URL or API key are empty
func New(url, apiKey string) (*Client, error) {
	if url == "" {
		return nil, errors.New("url is empty")
	}

	if apiKey == "" {
		return nil, errors.New("apiKey is empty")
	}

	return &Client{
		url:    url,
		apiKey: apiKey,
		client: &http.Client{},
	}, nil
}

// get sends a GET request to the specified endpoint and returns the response body as a byte slice.
// This function constructs the full URL by appending the endpoint to the base URL, sets up headers, and handles the HTTP response.
//
// Parameters:
//   - endpoint: API endpoint to send the GET request to
//
// Returns:
//   - []byte: Response body as a byte slice
//   - error: Error encountered during the request or response handling
func (c *Client) get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.url, endpoint)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", res.StatusCode, string(body))
	}

	return body, nil
}

// post sends a POST request with JSON data to the specified endpoint.
// This function marshals the given data into JSON, constructs the request, sets necessary headers,
// and processes the HTTP response.
//
// Parameters:
//   - endpoint: API endpoint to send the POST request to
//   - data: Data to be sent as JSON in the request body
//
// Returns:
//   - []byte: Response body as a byte slice
//   - error: Error encountered during the request or response handling
func (c *Client) post(endpoint string, data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshalling data to JSON: %w", err)
	}

	url := fmt.Sprintf("%s%s", c.url, endpoint)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", res.StatusCode, string(body))
	}

	return body, nil
}