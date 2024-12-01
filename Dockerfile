FROM golang:1.23-alpine AS builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o dating-sim-go .

FROM alpine:latest

WORKDIR /usr/local/bin

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/dating-sim-go .
COPY --from=builder /app/.env .

RUN chmod +x dating-sim-go

EXPOSE 5000

CMD ["dating-sim-go"]
