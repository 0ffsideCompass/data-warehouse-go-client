package client

import (
	"encoding/json"
	"fmt"
)

const (
	createPodcastEndpoint          = "/api/v1/podcasts"
	getPodcastEndpoint            = "/api/v1/podcasts/%s"
	searchPodcastsEndpoint        = "/api/v1/podcasts/search/%s"
	searchPodcastsPaginatedEndpoint = "/api/v1/podcasts/search/%s/paginated"
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
func (c *Client) CreatePodcast(request CreatePodcastRequest) (*Podcast, error) {
	responseData, err := c.post(createPodcastEndpoint, request)
	if err != nil {
		return nil, fmt.Errorf("error creating podcast: %w", err)
	}

	var response PodcastResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return response.Podcast, nil
}

// GetPodcast retrieves a specific podcast by its unique identifier.
//
// Parameters:
//   - id: A string representing the unique identifier of the podcast to retrieve
//
// Returns:
//   - *Podcast: The requested podcast
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON
func (c *Client) GetPodcast(id string) (*Podcast, error) {
	endpoint := fmt.Sprintf(getPodcastEndpoint, id)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error retrieving podcast: %w", err)
	}

	var response PodcastResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return response.Podcast, nil
}

// SearchPodcasts searches for podcasts that contain the specified tag.
//
// Parameters:
//   - tag: The tag to search for
//
// Returns:
//   - []Podcast: Array of podcasts matching the search criteria
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON
func (c *Client) SearchPodcasts(tag string) ([]Podcast, error) {
	endpoint := fmt.Sprintf(searchPodcastsEndpoint, tag)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error searching podcasts: %w", err)
	}

	var response PodcastsResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return response.Podcasts, nil
}

// SearchPodcastsPaginated searches for podcasts that contain the specified tag with pagination support.
// Page numbers start from 1. Limit must be between 1-100, defaults to 20.
//
// Parameters:
//   - tag: The tag to search for
//   - page: Page number (starts from 1)
//   - limit: Number of items per page (1-100)
//
// Returns:
//   - *PaginatedPodcastsResponse: Paginated response containing podcasts and pagination info
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON
func (c *Client) SearchPodcastsPaginated(tag string, page, limit int) (*PaginatedPodcastsResponse, error) {
	endpoint := fmt.Sprintf(searchPodcastsPaginatedEndpoint, tag)

	// Add query parameters
	endpoint = fmt.Sprintf("%s?page=%d&limit=%d", endpoint, page, limit)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error searching podcasts with pagination: %w", err)
	}

	var response PaginatedPodcastsResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return &response, nil
}