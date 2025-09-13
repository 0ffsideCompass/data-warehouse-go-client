# Data Warehouse Go Client

A Go client library for interacting with the Data Warehouse microservice API. This client provides a simple and efficient way to manage podcasts and articles with full CRUD operations, tag-based search, and pagination support.

## Features

- Full CRUD operations for Podcasts and Articles
- Tag-based search functionality
- Pagination support for large result sets
- Health check monitoring
- Bearer token authentication
- Comprehensive error handling
- Type-safe request and response models

## Installation

```bash
go get github.com/0ffsideCompass/data-warehouse-go-client
```

## Usage

### Initialize the Client

```go
package main

import (
    "log"
    client "github.com/0ffsideCompass/data-warehouse-go-client"
)

func main() {
    // Create a new client instance
    c, err := client.New("http://localhost:8080", "your-api-key")
    if err != nil {
        log.Fatal(err)
    }

    // Your code here...
}
```

### Health Check

```go
health, err := c.GetHealth()
if err != nil {
    log.Printf("Health check failed: %v", err)
    return
}

log.Printf("Service status: %s, Database: %s", health.Status, health.Database)
```

### Working with Podcasts

#### Create or Update a Podcast

```go
podcast, err := c.CreatePodcast(client.CreatePodcastRequest{
    Title:       "Tech Talk Podcast",
    Description: "Latest discussions on technology trends",
    URL:         "https://example.com/tech-talk",
    Tags:        []string{"technology", "programming", "innovation"},
})
if err != nil {
    log.Printf("Error creating podcast: %v", err)
    return
}

log.Printf("Podcast created with ID: %s", podcast.ID)
```

#### Get a Podcast by ID

```go
podcast, err := c.GetPodcast("507f1f77bcf86cd799439011")
if err != nil {
    log.Printf("Error retrieving podcast: %v", err)
    return
}

log.Printf("Retrieved podcast: %s", podcast.Title)
```

#### Search Podcasts by Tag

```go
// Simple search
podcasts, err := c.SearchPodcasts("technology")
if err != nil {
    log.Printf("Error searching podcasts: %v", err)
    return
}

for _, p := range podcasts {
    log.Printf("Found podcast: %s", p.Title)
}

// Paginated search
paginatedResult, err := c.SearchPodcastsPaginated("technology", 1, 20)
if err != nil {
    log.Printf("Error searching podcasts: %v", err)
    return
}

log.Printf("Found %d podcasts (Page %d of %d total)",
    len(paginatedResult.Podcasts),
    paginatedResult.Page,
    paginatedResult.Total)
```

### Working with Articles

#### Create or Update an Article

```go
article, err := c.CreateArticle(client.CreateArticleRequest{
    Title:       "Introduction to Clean Architecture",
    Description: "A comprehensive guide to implementing clean architecture in Go",
    URL:         "https://example.com/clean-architecture-guide",
    Tags:        []string{"architecture", "golang", "best-practices"},
})
if err != nil {
    log.Printf("Error creating article: %v", err)
    return
}

log.Printf("Article created with ID: %s", article.ID)
```

#### Get an Article by ID

```go
article, err := c.GetArticle("507f1f77bcf86cd799439012")
if err != nil {
    log.Printf("Error retrieving article: %v", err)
    return
}

log.Printf("Retrieved article: %s", article.Title)
```

#### Search Articles by Tag

```go
// Simple search
articles, err := c.SearchArticles("golang")
if err != nil {
    log.Printf("Error searching articles: %v", err)
    return
}

for _, a := range articles {
    log.Printf("Found article: %s", a.Title)
}

// Paginated search
paginatedResult, err := c.SearchArticlesPaginated("golang", 1, 20)
if err != nil {
    log.Printf("Error searching articles: %v", err)
    return
}

log.Printf("Found %d articles (Page %d of %d total)",
    len(paginatedResult.Articles),
    paginatedResult.Page,
    paginatedResult.Total)
```

## API Methods

### Podcast Methods

- `CreatePodcast(request CreatePodcastRequest) (*Podcast, error)` - Create or update a podcast
- `GetPodcast(id string) (*Podcast, error)` - Get a podcast by ID
- `SearchPodcasts(tag string) ([]Podcast, error)` - Search podcasts by tag
- `SearchPodcastsPaginated(tag string, page, limit int) (*PaginatedPodcastsResponse, error)` - Search podcasts with pagination

### Article Methods

- `CreateArticle(request CreateArticleRequest) (*Article, error)` - Create or update an article
- `GetArticle(id string) (*Article, error)` - Get an article by ID
- `SearchArticles(tag string) ([]Article, error)` - Search articles by tag
- `SearchArticlesPaginated(tag string, page, limit int) (*PaginatedArticlesResponse, error)` - Search articles with pagination

### Health Check

- `GetHealth() (*HealthResponse, error)` - Get service health status

## Error Handling

The client provides comprehensive error handling for various scenarios:

```go
podcast, err := c.GetPodcast("invalid-id")
if err != nil {
    // Error will contain details about the failure
    // e.g., "unexpected status code: 404, body: {"error":{"message":"podcast not found","code":"NOT_FOUND"}}"
    log.Printf("Error: %v", err)
}
```

## Authentication

The client uses Bearer token authentication. The API key is sent as a Bearer token in the Authorization header:

```
Authorization: Bearer your-api-key
```

## Types

### Request Types

- `CreatePodcastRequest` - Request payload for creating/updating podcasts
- `CreateArticleRequest` - Request payload for creating/updating articles

### Response Types

- `Podcast` - Podcast entity with all fields
- `Article` - Article entity with all fields
- `PodcastResponse` - Single podcast API response
- `ArticleResponse` - Single article API response
- `PodcastsResponse` - Multiple podcasts API response
- `ArticlesResponse` - Multiple articles API response
- `PaginatedPodcastsResponse` - Paginated podcasts with metadata
- `PaginatedArticlesResponse` - Paginated articles with metadata
- `HealthResponse` - Health check response
- `ErrorResponse` - API error response

## Development

### Running Tests

```bash
go test ./...
```

### Building

```bash
go build
```

## License

Private

## Support

For issues, questions, or contributions, please contact the Data Warehouse team.