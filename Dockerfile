FROM golang:1.22.3 as builder

WORKDIR /app
COPY go.mod ./
COPY . .
RUN go build -v -o main .

#########

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 9090

CMD ["./main"]
