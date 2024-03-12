FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

# Cloud build may not be installing golang
RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    golang && \
    rm -rf /var/lib/apt/lists/*

RUN go mod download

COPY . .

RUN go build -o bin/memecatcher ./cmd/api

CMD ["./bin/memecatcher"]