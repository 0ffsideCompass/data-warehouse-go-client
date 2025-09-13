package client

import (
	"encoding/json"
	"fmt"
)

const (
	createArticleEndpoint          = "/api/v1/articles"
	getArticleEndpoint            = "/api/v1/articles/%s"
	searchArticlesEndpoint        = "/api/v1/articles/search/%s"
	searchArticlesPaginatedEndpoint = "/api/v1/articles/search/%s/paginated"
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
func (c *Client) CreateArticle(request CreateArticleRequest) (*Article, error) {
	responseData, err := c.post(createArticleEndpoint, request)
	if err != nil {
		return nil, fmt.Errorf("error creating article: %w", err)
	}

	var response ArticleResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return response.Article, nil
}

// GetArticle retrieves a specific article by its unique identifier.
//
// Parameters:
//   - id: A string representing the unique identifier of the article to retrieve
//
// Returns:
//   - *Article: The requested article
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON
func (c *Client) GetArticle(id string) (*Article, error) {
	endpoint := fmt.Sprintf(getArticleEndpoint, id)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error retrieving article: %w", err)
	}

	var response ArticleResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return response.Article, nil
}

// SearchArticles searches for articles that contain the specified tag.
//
// Parameters:
//   - tag: The tag to search for
//
// Returns:
//   - []Article: Array of articles matching the search criteria
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON
func (c *Client) SearchArticles(tag string) ([]Article, error) {
	endpoint := fmt.Sprintf(searchArticlesEndpoint, tag)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error searching articles: %w", err)
	}

	var response ArticlesResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return response.Articles, nil
}

// SearchArticlesPaginated searches for articles that contain the specified tag with pagination support.
// Page numbers start from 1. Limit must be between 1-100, defaults to 20.
//
// Parameters:
//   - tag: The tag to search for
//   - page: Page number (starts from 1)
//   - limit: Number of items per page (1-100)
//
// Returns:
//   - *PaginatedArticlesResponse: Paginated response containing articles and pagination info
//   - error: An error object that reports issues either in sending the request, handling the response, or parsing the JSON
func (c *Client) SearchArticlesPaginated(tag string, page, limit int) (*PaginatedArticlesResponse, error) {
	endpoint := fmt.Sprintf(searchArticlesPaginatedEndpoint, tag)

	// Add query parameters
	endpoint = fmt.Sprintf("%s?page=%d&limit=%d", endpoint, page, limit)

	responseData, err := c.get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("error searching articles with pagination: %w", err)
	}

	var response PaginatedArticlesResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response data: %w", err)
	}

	return &response, nil
}