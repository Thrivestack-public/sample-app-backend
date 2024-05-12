FROM golang:1.18 as builder

WORKDIR /app
RUN go build -o app .

#########

FROM alpine:latest  

WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 9090

CMD ["./app"]
