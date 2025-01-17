FROM golang:1.23 as builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/api/main.go

RUN ls -l /app/server

FROM alpine:3.18

WORKDIR /root/

COPY --from=builder /app/server ./
COPY ./config/config.yaml ./config/config.yaml

RUN chmod +x ./server

CMD ["./server"]