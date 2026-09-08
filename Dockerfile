FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /bin/order ./cmd/api/order && \
  go build -o /bin/inventory ./cmd/api/inventory && \
  go build -o /bin/notification ./cmd/api/notification && \
  go build -o /bin/payment ./cmd/api/payment

FROM alpine:3

WORKDIR /app

COPY --from=builder /bin/order .
COPY --from=builder /bin/inventory  .
COPY --from=builder /bin/notification  .
COPY --from=builder /bin/payment .
