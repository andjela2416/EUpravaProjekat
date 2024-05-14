FROM golang:latest as builder
WORKDIR /app
COPY ./healthcare-service/go.mod ./healthcare-service/go.sum ./
RUN go mod download
COPY ./healthcare-service/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"]
