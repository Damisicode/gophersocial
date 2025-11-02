# ---- Build stage ----
FROM golang:1.24-alpine AS builder

# Install CA certificates inside builder (so we can copy them later)
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api cmd/api/*.go

# ---- Final stage ----
FROM scratch
WORKDIR /app

# Copy the CA bundle from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy your binary
COPY --from=builder /app/api .

EXPOSE 8080
CMD ["./api"]