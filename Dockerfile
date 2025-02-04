FROM golang:1.22-alpine3.18 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main -trimpath main.go

FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8082
CMD [ "/app/main" ]