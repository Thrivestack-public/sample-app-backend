FROM golang:1.18 as builder

WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

#########

FROM alpine:latest  

WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 9090

CMD ["./app"]
