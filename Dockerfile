# Create build stage based on buster image
FROM golang:1.16-buster AS builder
# Create working directory under /app
WORKDIR /app
# Copy over all go config (go.mod, go.sum etc.)
COPY ./rest ./rest
COPY go.* ./
# Install any required modules
RUN go mod download
# Copy over Go source code
COPY *.go ./
# Run the Go build and output binary under test
RUN go build -o /test
# Make sure to expose the port the HTTP server is using
EXPOSE 10000
# Run the app binary when we run the container
ENTRYPOINT ["/test"]

# docker build -t test_go_app .
# docker run -p 10000:10000 -t test_go_app