package client

import (
	"fmt"

	"github.com/0ffsideCompass/models"
)

const (
	createArticleEndpoint = "/api/v1/articles"
)

// CreateArticle creates or updates an article in the Data Warehouse.
// If an article with the same URL already exists, it will be updated.
//
// Parameters:
//   - request: CreateArticleRequest containing the article details
//
// Returns:
//   - *Article: The created or updated article
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON
func (c *Client) CreateArticle(request models.DataWarehouseCreateArticleRequest) error {
	_, err := c.post(createArticleEndpoint, request)
	if err != nil {
		return fmt.Errorf("error creating article: %w", err)
	}

	return nil
}
