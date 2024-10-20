FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o medilearn-backend .

# Final image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/medilearn-backend .

EXPOSE 3000

CMD ["./medilearn-backend"]

