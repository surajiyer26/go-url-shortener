# Go URL Shortener

A simple URL shortener service written in Go. This project provides a basic web API to shorten URLs and redirect users to the original URLs. It supports in-memory and Redis-based storage backends.

## Features
- Shorten long URLs to short codes
- Redirect short codes to original URLs
- Pluggable storage: in-memory and Redis

## Project Structure
```
go.mod                # Go module file
main.go               # Application entry point
handlers/             # HTTP handlers
  redirect.go         # Handles redirection from short URL to original URL
  shorten.go          # Handles URL shortening requests
storage/              # Storage backends
  memory_store.go     # In-memory storage implementation
  redis_store.go      # Redis storage implementation
  store.go            # Storage interface definition
```

## Getting Started
1. **Clone the repository**
   ```
   git clone https://github.com/surajiyer26/go-url-shortener.git
   cd go-url-shortener
   ```
2. **Install dependencies**
   ```
   go mod tidy
   ```
3. **Run the application**
   ```
   go run main.go
   ```

## API Endpoints
- `POST /shorten` - Shorten a URL
- `GET /{shortCode}` - Redirect to the original URL

## Storage Options
- **In-memory**: Default, no setup required.
- **Redis**: Configure Redis connection by setting `REDIS_ADDR` env var.

## License
MIT
