FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /bin/order ./cmd/order && \
  go build -o /bin/inventory ./cmd/inventory && \
  go build -o /bin/notification ./cmd/notification && \
  go build -o /bin/payment ./cmd/payment

FROM alpine:3

WORKDIR /app

COPY --from=builder /bin/order .
COPY --from=builder /bin/inventory  .
COPY --from=builder /bin/notification  .
COPY --from=builder /bin/payment .
