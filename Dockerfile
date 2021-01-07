FROM golang:1.11.1 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./...

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]