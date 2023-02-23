# Build stage
FROM golang:1.19.1-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration
RUN chmod +x start.sh
RUN chmod +x wait-for.sh

EXPOSE 8085
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]