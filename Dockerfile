FROM golang:alpine
WORKDIR /app
RUN apk update --no-cache
COPY . .
RUN go build
CMD ["./v2-Personal-Website"]

