# URL Shortener in Go

A simple URL shortener service built with Go

[url-shortener.itskoshkin.site](https://url-shortener.itskoshkin.site)

## Features

*   URL Shortening – сonvert long URLs into short, easy-to-share links
*   User Link Management – tracks user links using a persistent cookie-based ID
*   ID Portability – users can export and import their ID to use service across different browsers
*   Admin Panel – password-protected page to view and manage all existing links

## Getting Started

### Prerequisites

*   Go (version 1.20+ recommended)
*   PostgreSQL server running locally or remotely

### Installation & Setup

1.  **Clone the repository**
    ```bash
    git clone https://github.com/itsk-example-projects/url-shortener.git
    cd url-shortener
    ```

2.  **Configure the application**

    Edit `config.yaml` file in the root directory of the project and set your domain, database and admin credentials<br><br>

3.  **Install dependencies**
    ```bash
    go mod tidy
    ```

4.  **Run the application**
    ```bash
    go run cmd/main.go
    ```
    The server will start on [`http://localhost:8080`](http://localhost:8080) by default, you will need to use `nginx` or another reverse proxy to serve it via domain<br>
    Database must be created manually, table `links` will be created automatically on first run
