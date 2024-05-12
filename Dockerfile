FROM golang:1.22 as builder

WORKDIR /app
COPY go.mod ./
RUN go build -o app .

#########

FROM alpine:latest  

WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 9090

CMD ["./app"]
