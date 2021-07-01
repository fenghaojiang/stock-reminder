FROM golang:1.16 AS builder

WORKDIR /go/src/app
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go get -d -v ./... \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o stock_reminder .


FROM alpine:latest
RUN apk update \
    && apk add tzdata \
    && apk --no-cache add ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo 'Asia/Shanghai' >/etc/timezone 
WORKDIR /root/
COPY --from=builder /go/src/app/stock_reminder .

CMD ["./stock_reminder"]


