FROM golang:alpine
WORKDIR /app
RUN apk update --no-cache
COPY . .
ENV GITHUB_USERNAME="6ixBit"
ENV GITHUB_URL="https://api.github.com/users/6ixbit/repos?direction=desc"
RUN go build
CMD ["./v2-Personal-Website"]

