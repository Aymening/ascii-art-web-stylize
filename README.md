# Ascii-art-web

Ascii-art-web-stylize is a Go-based web server project designed to enhance user experience and visual appeal while generating ASCII art from text input. It provides an interactive and intuitive interface for users to stylize their ASCII art using various banners (shadow, standard, thinkertoy).

## Authors
- [alahmami]
- [bbenskou]
- [ahiti]

## Usage

### How to Run
To run the Ascii-art-web-stylize, follow these steps:

1. Clone the repository: `git clone https://learn.zone01oujda.ma/git/bbenskou/ascii-art-web-stylize`
2. Navigate to the project directory: `cd ascii-art-web-stylize`
3. Run the server: `go run main.go`

### Accessing the Web Interface
Once the server is running, access the web interface by opening a web browser and navigating to `http://localhost:8080`.

## Implementation Details

### Algorithm
The Ascii-art-web-stylize project implements the following algorithm:

1. Set up a Go HTTP server to handle incoming requests.
2. Define HTTP endpoints:
   - `GET /`: Sends HTML response, the main page with input fields and options to select banners.
   - `POST /ascii-art`: Receives data from the client (text and selected banner) and generates ASCII art.
3. Utilize Go templates to render HTML pages and display data dynamically.
4. Handle HTTP status codes appropriately:
   - OK (200) for successful requests.
   - Not Found (404) for missing resources.
   - Bad Request (400) for incorrect requests.
   - Internal Server Error (500) for unhandled errors.