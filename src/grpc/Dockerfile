FROM golang:1.24.0-bookworm

WORKDIR /go/src

RUN apt update && \
    apt install -y protobuf-compiler

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download
