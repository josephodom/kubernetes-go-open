# syntax=docker/dockerfile:1
FROM golang:1.23-alpine
WORKDIR /

COPY src/ .

RUN go mod init github.com/josephodom/kubernetes-go-open
RUN go get github.com/redis/go-redis/v9

RUN go build -o /bin/main main.go

FROM golang:1.23-alpine
COPY --from=0 /bin/main /bin/main
CMD ["/bin/main"]
