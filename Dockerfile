FROM golang:1.13 AS builder
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/main .
EXPOSE 3000
ENTRYPOINT ["./main"]