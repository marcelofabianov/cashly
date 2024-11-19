FROM golang:1.23-bookworm

ENV PATH=$PATH:/root/go/bin

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apt-get update && apt-get install -y \
    bash \
    unzip \
    wget \
    openssl && \
    wget -q https://github.com/protocolbuffers/protobuf/releases/download/v28.3/protoc-28.3-linux-x86_64.zip -O protoc.zip && \
    unzip protoc.zip -d /usr/local && \
    rm protoc.zip && \
    apt-get clean && rm -rf /var/lib/apt/lists/* && \
    go install github.com/pressly/goose/v3/cmd/goose@latest && \
    go install github.com/air-verse/air@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN chmod +x _scripts/generate/*.sh

RUN go build -o /app/main/api ./cmd/api/main.go

EXPOSE 50051

CMD ["air", "-c", "/app/.air.toml"]
