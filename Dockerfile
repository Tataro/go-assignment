FROM golang:1.10

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/gitlab.com/upaphong/go-assignment
COPY . ./
RUN dep ensure --vendor-only