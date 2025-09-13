package client

import (
	"encoding/json"
	"fmt"
)

const (
	healthEndpoint = "/health"
)

// GetHealth retrieves the health status of the Data Warehouse service.
// This endpoint is useful for monitoring and health checks.
//
// Returns:
//   - *HealthResponse: The health status of the service
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON
func (c *Client) GetHealth() (*HealthResponse, error) {
	responseData, err := c.get(healthEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error retrieving health status: %w", err)
	}

	var response HealthResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return &response, nil
}