FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main ./cmd/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

COPY ./db/tickets.csv ./db/

EXPOSE 8080

CMD ["./main"]
