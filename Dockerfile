# Build stage
FROM golang:1.19.1-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go
RUN apk --no-cache add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8085
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]