FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /bin/order ./cmd/order/api && \
  go build -o /bin/inventory ./cmd/inventory/api && \
  go build -o /bin/notification ./cmd/notification/api && \
  go build -o /bin/payment ./cmd/payment/api

FROM alpine:3

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /bin/order .
COPY --from=builder /bin/inventory  .
COPY --from=builder /bin/notification  .
COPY --from=builder /bin/payment .
