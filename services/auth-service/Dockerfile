# --- Stage 1: Build ---
    FROM golang:1.23-alpine AS builder

    # Install git/ca-certificates if your Go app needs to fetch private modules or talk to HTTPS APIs
    RUN apk add --no-cache git ca-certificates
    
    WORKDIR /app
    
    # Copy only dependency files first to leverage Docker layer caching
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy the rest of the source code
    COPY . .
    
    # Build the binary
    # CGO_ENABLED=0 creates a statically linked binary that runs on 'scratch'
    RUN CGO_ENABLED=0 GOOS=linux go build -o /krakit-binary ./cmd/main.go
    
    # --- Stage 2: Final Runtime ---
    FROM scratch
    
    # Import CA certificates from builder to allow HTTPS calls
    COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
    
    # Copy the compiled binary from the builder stage
    COPY --from=builder /krakit-binary /krakit-binary
    
    # Ensure the binary is executable
    ENTRYPOINT ["/krakit-binary"]