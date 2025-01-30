# Step 1: Use the official Golang image to build the Go application
FROM golang:1.22.2-alpine AS builder

# Step 2: Set the working directory inside the container
WORKDIR /app

# Step 3: Copy go mod and sum files to download dependencies
COPY go.mod go.sum ./

# Step 4: Download Go module dependencies
RUN go mod tidy

# Step 5: Copy the entire project code into the container
COPY . .

# Step 6: Build the Go application
RUN go build -o gomongo ./cmd/server

# Step 7: Create a new stage to run the application (this avoids shipping the full Go SDK)
FROM alpine:latest

# Step 8: Set the working directory inside the container
WORKDIR /root/

# Step 9: Copy the built executable from the builder stage
COPY --from=builder /app/gomongo .

# Step 10: Expose the application port
EXPOSE 8080

# Step 11: Command to run the application
CMD ["./gomongo"]
