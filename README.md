# v2-Personal-Website
My personal website rewritten from a full stack Python app to a Golang backend API

## Features
- Request Logging
- Concurrent email sending
- Caching
- Rate limiting
- Monitoring
- Containerised
- COnfigurable via environment
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

### Kubernetes access
In order to use the kubernetes deployment and the service so that it can be accessed do the following;

- kubectl apply -f deployBackend.yml

To access the application, get the INTERNAL-IP address of the Node and prefix the approparite port to it (xxx.xxx.xxx.xxx:30500)

- kubectl get nodes

Then you can visit the endpoints for the application
