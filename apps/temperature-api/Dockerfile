FROM golang:1.21-alpine as builder
WORKDIR /app
COPY main.go .
RUN go build -o temperature-api main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/temperature-api .
EXPOSE 8081
ENTRYPOINT ["./temperature-api"]
