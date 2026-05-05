FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o realtime ./cmd/server

FROM scratch
COPY --from=builder /app/realtime /realtime
EXPOSE 9000
ENTRYPOINT ["/realtime"]
