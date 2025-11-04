FROM golang:1.25

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o task-api ./cmd/main.go

EXPOSE 8080
CMD ["./task-api"]