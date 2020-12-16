# v2-Personal-Website
My personal website rewritten from a full stack Python app to a Golang backend API

## Features
- Request Logging
- Concurrent email sending
- Caching
- Containerised
***

### Run
##### Binary executable
- go build
- go run ./v2-Personal-Website
- visit localhost:8080
***

##### Docker container
- docker build -t mywebserver:latest .
- docker run -p 8080:8080 mywebserver:latest
- visit localhost:8080