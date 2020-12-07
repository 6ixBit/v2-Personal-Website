# v2-Personal-Website
My personal website rewritten and updated in Golang

## Features
- Request Logging
- Concurrent email sending
- Caching
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