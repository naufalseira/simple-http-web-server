# Simple HTTP Web Server

A simple HTTP web server built with Go and the Gorilla Mux framework, designed to demonstrate core web server functionalities including static file serving, static routing, dynamic routing, form handling, and middleware for HTTP method validation and logging.

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
  - **Logging Middleware**: Logs request details like the HTTP method, accessed URL and processing time.
  - **HTTP Method Middleware**: Validates allowed HTTP methods for each endpoint.

## Requirements

- **Go 1.22.0**: Ensure Go is installed on your system.

## How to Use

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
5. The server will run at:
```bash
  http://localhost:8090
  ```

## Running with Docker

To run this project using Docker, you can pull the Docker image from Docker Hub and start a container.
1. Pull the Docker Image:
```bash
  docker pull 0xsera/simple-http-web-server
  ```
2. Start a container from the image:
```bash
  docker run -p 8090:8090 0xsera/simple-http-web-server
  ```
3. With the container running, you can access the application in your browser at:
```bash
  http://localhost:8090
  ```

## Usage Examples
- http://localhost:8090/ : Displays the main page.
- http://localhost:8090/#contact : Submits a form with data such as name, email, and message.
- http://localhost:8090/jakarta : Displays the welcome message for the Jakarta page.
- http://localhost:8090/destination/Bali : Displays the welcome message for the Bali page.