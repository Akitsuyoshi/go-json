FROM golang:1.8.3

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/Akitsuyoshi/go-json/

# Install our dependencies
RUN go get github.com/gorilla/mux

# Install api binary globally within container
RUN go install github.com/Akitsuyoshi/go-json/api

# Set binary as entrypoint
ENTRYPOINT /go/bin/api

# Expose default port (8081)
EXPOSE 8081
