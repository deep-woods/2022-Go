# Golang image
FROM golang:1.17 AS builder

ENV GO111MODULE=on

# Create a directory '/app' as the default workspace inside a container using this image. 
WORKDIR /app  
ADD . /app
RUN go get -u github.com/githubnemo/CompileDaemon
 
# In Golang, Dependencies are managed by go.mod.
# Equivalent: package.json in nodejs. 

# The following two lines are environment setup for the Docker container using this image. 
# Copy go.mod and go.sum into the image so they take care of the versions and dependencies inside future containers. 
# COPY <src> <dest>: The <dest> is an absolute path, or a path relative to WORKDIR
# COPY go.mod and go.sum to WORKDIR (/app)
COPY go.mod .
COPY go.sum . 

# The go mod download command downloads the named modules into the module cache. 
# Arguments can be module paths or module patterns selecting dependencies 
# of the main module or version queries of the form path@version. 
# With no arguments, download applies to all dependencies of the main module.
# The go command will automatically download modules as needed during ordinary execution. 
# The go mod download command is useful mainly for pre-filling the module cache 
# or for loading data to be served by a module proxy.
RUN go mod download
COPY . .
 
# golang is a compile language - build is required.
ENTRYPOINT CompileDaemon --polling --log-prefix=false --build='go build main.go' --command=./main