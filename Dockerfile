FROM golang:1.8.3-alpine

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/Akitsuyoshi/go-json/api

# Install our dependencies

# Install api binary globally within container
RUN go install github.com/Akitsuyoshi/go-json/api

# Set binary as entrypoint
ENTRYPOINT /go/bin/api

# Expose default port (80)
EXPOSE 80
