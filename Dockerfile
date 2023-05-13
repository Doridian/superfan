FROM golang:1.20-alpine AS builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

COPY . /app
RUN go build -o /superfan ./cmd/superfan

FROM alpine:3.18
COPY LICENSE /LICENSE

RUN apk add --no-cache lm-sensors

COPY --from=builder --chown=0:0 --chmod=755 /superfan /superfan

ENTRYPOINT ["/superfan"]
