FROM golang:1.22 as builder

WORKDIR /app
COPY go.mod ./
COPY . .
RUN go build -v -o app .

#########

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 9090

CMD ["./app"]
