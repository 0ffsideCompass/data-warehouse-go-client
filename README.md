# Data Warehouse Go Client

A Go client library for interacting with the OffsideCompass Data Warehouse microservice. This client provides a simple and secure interface for managing articles and podcasts in the data warehouse.

## Installation

```bash
go get github.com/0ffsideCompass/data-warehouse-go-client
```

## Features

-  Simple HTTP client with built-in authentication
-  Support for creating/updating articles and podcasts
-  Automatic API key authentication
-  JSON request/response handling
-  Comprehensive error handling

## Usage

### Initialize the Client

```go
import "github.com/0ffsideCompass/data-warehouse-go-client/client"

// Create a new client instance
dwClient, err := client.New("https://api.example.com", "your-api-key")
if err != nil {
    log.Fatal("Failed to create client:", err)
}
```

### Create or Update an Article

```go
import "github.com/0ffsideCompass/models"

articleRequest := models.DataWarehouseCreateArticleRequest{
    Title:       "Article Title",
    URL:         "https://example.com/article",
    Content:     "Article content...",
    Author:      "John Doe",
    PublishedAt: time.Now(),
    // Add other fields as needed
}

err := dwClient.CreateArticle(articleRequest)
if err != nil {
    log.Printf("Failed to create article: %v", err)
}
```

### Create or Update a Podcast

```go
import "github.com/0ffsideCompass/models"

podcastRequest := models.DataWarehouseCreatePodcastRequest{
    Title:       "Podcast Title",
    URL:         "https://example.com/podcast",
    Description: "Podcast description...",
    Duration:    3600, // in seconds
    PublishedAt: time.Now(),
    // Add other fields as needed
}

err := dwClient.CreatePodcast(podcastRequest)
if err != nil {
    log.Printf("Failed to create podcast: %v", err)
}
```

## API Reference

### Client Methods

#### `New(url, apiKey string) (*Client, error)`
Creates a new client instance with the specified base URL and API key.

#### `CreateArticle(request models.DataWarehouseCreateArticleRequest) error`
Creates or updates an article in the Data Warehouse. If an article with the same URL already exists, it will be updated.

#### `CreatePodcast(request models.DataWarehouseCreatePodcastRequest) error`
Creates or updates a podcast in the Data Warehouse. If a podcast with the same URL already exists, it will be updated.

## Error Handling

The client provides detailed error messages for various failure scenarios:

- Empty URL or API key during client initialization
- Network connectivity issues
- Authentication failures
- Invalid request data
- Server errors

All errors are wrapped with context to help with debugging.

## Requirements

- Go 1.24.1 or higher
- Dependencies:
  - `github.com/0ffsideCompass/models` v1.0.2
  - `go.mongodb.org/mongo-driver` v1.17.1 (indirect)

## Security

The client uses Bearer token authentication. Ensure your API key is kept secure and never committed to version control.

## Support

[Add support contact information if applicable]