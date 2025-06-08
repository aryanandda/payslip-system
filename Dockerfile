# Start from the official Go image
FROM golang:1.23 AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    TZ=Asia/Jakarta

# Create working directory
WORKDIR /app

# Copy Go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o payslip-system ./cmd/main.go

# Final stage - clean image
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone

WORKDIR /root/
COPY --from=builder /app/payslip-system .

EXPOSE 8080

CMD ["./payslip-system"]
