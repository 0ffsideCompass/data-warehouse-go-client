package client

import (
	"fmt"

	"github.com/0ffsideCompass/models"
)

const (
	createPodcastEndpoint = "/api/v1/podcasts"
)

// CreatePodcast creates or updates a podcast in the Data Warehouse.
// If a podcast with the same URL already exists, it will be updated.
//
// Parameters:
//   - request: CreatePodcastRequest containing the podcast details
//
// Returns:
//   - *Podcast: The created or updated podcast
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON
func (c *Client) CreatePodcast(request models.DataWarehouseCreatePodcastRequest) error {
	_, err := c.post(createPodcastEndpoint, request)
	if err != nil {
		return fmt.Errorf("error creating podcast: %w", err)
	}

	returnnil
}
