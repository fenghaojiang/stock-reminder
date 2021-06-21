FROM golang:1.16

WORKDIR /go/src/app
COPY . .

RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo 'Asia/Shanghai' >/etc/timezone \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go get -d -v ./... \
    && go install -v ./...


CMD ["stock_reminder"]


