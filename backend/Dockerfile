# Build stage
FROM golang:1.24 AS build

WORKDIR /backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app ./cmd/quiz/quiz.go

# Final minimal image
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add dumb-init ca-certificates

COPY --from=build /backend/app .

RUN chmod +x ./app

EXPOSE 3000

ENTRYPOINT ["/usr/bin/dumb-init", "--", "./app"]
