FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o main cmd/api/main.go

FROM alpine:3.18

ARG PORT

ENV PORT=${PORT}

EXPOSE ${PORT}

COPY --from=builder /app/main /app/main

CMD ["./app/main"]
