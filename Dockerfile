# Step 1: Use Golang Alpine image with CGO dependencies
FROM golang:1.22-alpine AS builder

# Step 2: Install build dependencies (gcc and musl-dev)
RUN apk add --no-cache gcc musl-dev

# Step 3: Set working directory and enable CGO
WORKDIR /app
# Allow Go to call C libraries
ENV CGO_ENABLED=1

# Step 4: Copy dependency files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 5: Copy the source code and build the binary
# Copies all source code and files from local directory to the current 
# working directory in the container
COPY . .

# Compiles the Go code and creates a binary file named gopportunities-api.
# This binary will be placed in the /app directory.
RUN go build -o gopportunities-api ./cmd

# Step 6: Create a minimal final image
FROM alpine:latest

# Step 7: Install SQLite libraries on the final image
RUN apk add --no-cache sqlite-libs

# Step 8: Set working directory and copy the binary from the builder
WORKDIR /root/
COPY --from=builder /app/gopportunities-api .

# Step 9: Expose port and run the API
EXPOSE 8080
# Specifies the command to run when the container starts. 
# It runs the gopportunities-api binary
CMD ["./gopportunities-api"]
