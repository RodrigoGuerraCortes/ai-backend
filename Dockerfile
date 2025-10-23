# ---------- Stage 1: Build ----------
FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod & go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o ai-backend ./cmd/api/main.go

# ---------- Stage 2: Runtime ----------
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/ai-backend .
#COPY .env .env

EXPOSE 8080

CMD ["./ai-backend"]
