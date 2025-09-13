package client

import "time"

// Podcast represents a podcast entity in the Data Warehouse
type Podcast struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url"`
	Tags        []string  `json:"tags,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Article represents an article entity in the Data Warehouse
type Article struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url"`
	Tags        []string  `json:"tags,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreatePodcastRequest represents the request payload for creating or updating a podcast
type CreatePodcastRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	URL         string   `json:"url"`
	Tags        []string `json:"tags,omitempty"`
}

// CreateArticleRequest represents the request payload for creating or updating an article
type CreateArticleRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description,omitempty"`
	URL         string   `json:"url"`
	Tags        []string `json:"tags,omitempty"`
}

// PodcastResponse represents a single podcast response from the API
type PodcastResponse struct {
	Podcast *Podcast `json:"podcast"`
}

// ArticleResponse represents a single article response from the API
type ArticleResponse struct {
	Article *Article `json:"article"`
}

// PodcastsResponse represents multiple podcasts response from the API
type PodcastsResponse struct {
	Podcasts []Podcast `json:"podcasts"`
}

// ArticlesResponse represents multiple articles response from the API
type ArticlesResponse struct {
	Articles []Article `json:"articles"`
}

// PaginatedPodcastsResponse represents a paginated response for podcasts
type PaginatedPodcastsResponse struct {
	Podcasts []Podcast `json:"podcasts"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	Limit    int       `json:"limit"`
}

// PaginatedArticlesResponse represents a paginated response for articles
type PaginatedArticlesResponse struct {
	Articles []Article `json:"articles"`
	Total    int       `json:"total"`
	Page     int       `json:"page"`
	Limit    int       `json:"limit"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Database  string    `json:"database"`
	Timestamp time.Time `json:"timestamp"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"error"`
}