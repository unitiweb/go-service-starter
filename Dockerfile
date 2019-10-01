# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang:1.12-alpine base image
FROM golang:1.12-alpine

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

RUN mkdir -p /do/github.com/unitiweb/go-microservice

# Set the Current Working Directory inside the container
WORKDIR /do/github.com/unitiweb/go-microservice

# Copy go mod and sum files
#COPY go.mod go.sum ./
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . /do/github.com/unitiweb/go-microservice

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
#EXPOSE 8080

# Run the executable
#CMD ["./main/main"]