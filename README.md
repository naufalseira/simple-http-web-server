# Simple HTTP Web Server

A simple web application built with Go, using the Gorilla Mux framework to create an HTTP server. The application serves both static and dynamic endpoints and includes middleware for logging and HTTP method validation.

## Features

- **Static File Server**: Serves static files from the `/assets` directory.
- **Static Endpoints**:
  - `/`: Displays the main `index.html` page.
  - `/jakarta`: Shows a welcome message for the Jakarta page.
- **Dynamic Endpoints**:
  - `/destination/{place}`: Displays a welcome message for a specific destination provided in the URL.
- **Form Handling**:
  - `/contact`: Accepts POST requests to process submitted form data.
- **Middleware**:
  - **Logging Middleware**: Logs the HTTP method, accessed URL and processing time.
  - **HTTP Method Middleware**: Validates allowed HTTP methods for each endpoint.

## Requirements

- **Go 1.22.0**: Ensure Go is installed on your system.

## Installation

1. Clone this repository:
  ```bash
  git clone https://github.com/naufalseira/simple-http-web-server.git
  ```
2. Change directory to main folder:
```bash
  cd simple-http-web-server
  ```
3. Download all depedency:
```bash
  go mod download
  ```
4. Start the server:
```bash
  go run main.go
  ```
5. The server will run at http://localhost:8090. Available endpoints:
- http://localhost:8090/
- http://localhost:8090/#contact
- http://localhost:8090/jakarta
- http://localhost:8090/destination/{place}

## Usage Examples
- http://localhost:8090/ : Displays the main page.
- http://localhost:8090/#contact : Submits a form with data such as name, email, and message.
- http://localhost:8090/jakarta : Displays the welcome message for the Jakarta page.
- http://localhost:8090/destination/Bali : Displays the welcome message for the Bali page.