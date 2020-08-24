# Build the cipher service binary
FROM golang:1.12.7-stretch as base

# add the working directory for the project
WORKDIR /go/src/watermarksvc

# Copy the service code
COPY api api
COPY cmd/authorization cmd/authorization
COPY internal internal
COPY pkg/authorization pkg/authorization
COPY vendor vendor
COPY go.mod go.mod
COPY go.sum go.sum

# building service binary at path discovergy/www
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOFLAGS=-mod=vendor -o auth go build ./cmd/authorization

## Using the multi-stage image to just run the binary
FROM alpine:latest as final

# Init working directory to root /
WORKDIR /

# Copy just the binary from the base image
COPY --from=base /go/src/watermarksvc/auth .

# just an indication that this port will be exposed by this container
EXPOSE 3335

# command to run at the immediate start of the container
ENTRYPOINT ["./auth"]
