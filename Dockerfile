# syntax=docker/dockerfile:1
FROM golang:1.23
WORKDIR /

COPY src/ .

# RUN go mod init github.com/josephodom/kubernetes-go-open
COPY go.mod .

RUN go build -o /bin/main main.go

FROM scratch
COPY --from=0 /bin/main /bin/main
CMD ["/bin/main"]
