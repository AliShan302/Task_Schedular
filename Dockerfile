FROM golang:1.18.1-alpine3.14
WORKDIR /app
COPY . .
RUN go mod tidy
EXPOSE 8080
CMD ["go", "run", "cmd/task-schedular-server/main.go"]
