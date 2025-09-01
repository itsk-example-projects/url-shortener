# URL Shortener in Go

A simple URL shortener service built with Go

[url-shortener.itskoshkin.site](https://url-shortener.itskoshkin.site)

## Features

- URL Shortening – сonvert long URLs into short, easy-to-share links
- User Link Management – tracks user links using a persistent cookie-based ID
- ID Portability – users can export and import their ID to use service across different browsers
- Admin Panel – password-protected page to view and manage all existing links

## Getting Started

### Prerequisites

- Go 1.23+
- PostgreSQL server running locally or remotely

### Running
1. Plain
    1. Clone the repository
        ```bash
        git clone https://github.com/itsk-example-projects/url-shortener.git
        cd url-shortener
        ```
    2. Install dependencies
        ```bash
        go mod tidy
        ```
    3. Run the application
        ```bash
        go run cmd/main.go
        ```
       *(Use `screen` to run app in a separate session so your terminal stays free)*
2. Docker
    1. Clone the repository
        ```bash
        git clone https://github.com/itsk-example-projects/url-shortener.git
        cd url-shortener
        ```
    2. Modify config  
        Edit `config.yaml` to change Postgres connection params and admin password
    3. Build Docker image
        ```bash
        docker build -t url-shortener-example:latest .
        ```
    4. Run the container
        ```bash
        docker run -dit --name url-shortener-example -p 8080:8080 url-shortener-example:latest
        ```
       *App now available at http://localhost:8080*