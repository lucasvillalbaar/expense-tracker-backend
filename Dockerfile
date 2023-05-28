ARG GO_VERSION=1.20.4

FROM golang:${GO_VERSION}-alpine AS builder

RUN go env -w GOPROXY=direct
RUN apk add --no-cache git
RUN apk --no-cache add ca-certificates && update-ca-certificates

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY database database
COPY expense-tracker-service expense-tracker-service 
COPY internal internal
COPY pkg pkg
COPY repository repository

RUN go install ./...

FROM alpine:latest
WORKDIR /usr/bin

COPY --from=builder /go/bin .
