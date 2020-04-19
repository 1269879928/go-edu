FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /go/go-edu

WORKDIR /go/go-edu

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7

ENV REDIS_ADDR=""
ENV REDIS_PW=""
ENV REDIS_DB=""
ENV MysqlDSN="mysql57:shijinting0510@tcp(127.0.0.1:3306)/edu?charset=utf8mb4&parseTime=True&loc=Local"
ENV GIN_MODE="release"
ENV PORT=3000

RUN apk update && \
    apk add ca-certificates

COPY --from=build /go/go-edu/api_server /usr/bin/api_server

RUN chmod +x /usr/bin/api_server
CMD /usr/sbin/init
EXPOSE 3000
ENTRYPOINT ["api_server"]
