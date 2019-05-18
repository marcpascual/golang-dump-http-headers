# docker build -t go-http-dump-headers .
# docker run -d --rm -p 3000:3000 go-dump-http-headers
FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go build server.go

CMD ["/go/src/app/server"]
EXPOSE 3000
