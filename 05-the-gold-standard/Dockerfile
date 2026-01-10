# LEVEL 5: THE GOLD STANDARD (HARDENING MAX)
# This example uses Go (Compiled) + Distroless (Static).
# There is no shell (/bin/sh), no 'ls', no 'cd'. Nothing but the binary.

# --- STAGE 1: Build ---
# Using a specific version tag for reproducibility
FROM golang:1.21-alpine AS builder

# Install SSL certs (needed for HTTPS calls from binary)
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY main.go .

# Build as a static binary (no dependencies on system libraries)
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# --- STAGE 2: Hardening ---
# Google's Distroless image contains ONLY the app and its runtime dependencies.
# It does NOT contain a shell, which disables 99% of remote execution exploits.
FROM gcr.io/distroless/static-debian12

WORKDIR /

# 1. Copy the binary
COPY --from=builder /app/main /main
# 2. Copy SSL certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# 3. Security: Run as a non-root UID (Distroless 'nonroot' user is 65532)
USER 65532:65532

EXPOSE 8080

# Since there is no shell, we must use the Exec (JSON) form
ENTRYPOINT ["/main"]
