FROM golang:1.25-alpine

WORKDIR /app

# Install Air for hot reloading
RUN go install github.com/air-verse/air@v1.61.7

# Copy dependency file and download them
COPY go.mod go.su[m] ./
RUN go mod download

# Start Air
CMD [ "air", "--build.cmd", "go build -o tmp/main ./cmd", "--build.bin", "./tmp/main" ]