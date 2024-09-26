FROM golang:1.23-bullseye AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y \
    libtesseract-dev \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -o main .



FROM debian:bullseye-slim AS webui

RUN apt-get update && apt-get install -y \
    libleptonica-dev \
    libtesseract-dev \
    tesseract-ocr \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
