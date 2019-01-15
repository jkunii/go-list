FROM golang:1.11.1-stretch

LABEL maintainer "fukanoki@gmail.com"

RUN curl https://glide.sh/get | sh




ENV PKG_NAME=github.com/jkunii/go-list
ENV PKG_PATH=$GOPATH/src/$PKG_NAME
WORKDIR $PKG_PATH

COPY glide.yaml glide.lock $PKG_PATH/
RUN glide install

COPY . $PKG_PATH
RUN go build && go install

WORKDIR $PKG_PATH
EXPOSE 1323
CMD ["go-list"]