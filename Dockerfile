FROM golang:1.18 as builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go build -o app .

#########

FROM alpine:latest  

WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 9090

CMD ["./app"]
