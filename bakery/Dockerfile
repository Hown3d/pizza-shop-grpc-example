# STEP 1 build executable binary
############################
FROM golang:1.16-buster AS builder
WORKDIR /app
COPY go.mod go.sum ./
# Fetch dependencies.
# Using go mod download.
RUN go mod download 

COPY . .

# Build the binary.
RUN go build -o ./bakery
############################
# STEP 2 build a small image
############################
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy our static executable.
COPY --from=builder /app/bakery /root/bakery
# Run the hello binary.
CMD ["/root/bakery"]