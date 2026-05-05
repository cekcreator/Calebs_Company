FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o auth ./cmd/server

FROM scratch
COPY --from=builder /app/auth /auth
EXPOSE 8001
ENTRYPOINT ["/auth"]
