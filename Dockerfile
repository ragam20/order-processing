FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/order ./cmd/order/api && \
    CGO_ENABLED=0 go build -o /bin/inventory ./cmd/inventory/api && \
    CGO_ENABLED=0 go build -o /bin/notification ./cmd/notification/api && \
    CGO_ENABLED=0 go build -o /bin/payment ./cmd/payment/api

FROM alpine:3

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /bin/order /bin/order
COPY --from=builder /bin/inventory /bin/inventory
COPY --from=builder /bin/notification /bin/notification
COPY --from=builder /bin/payment /bin/payment
