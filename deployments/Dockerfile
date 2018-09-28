FROM golang:1.10-alpine as builder


ENV PKG_NAME=crypto-util
ENV PKG_SCOPE=github.com/edimarlnx/go-crypto-util

COPY ./ /go/src/${PKG_SCOPE}

WORKDIR /go/src/${PKG_SCOPE}/pkg/crypto
RUN go install --ldflags '-s -w -extldflags "-static"'

WORKDIR /go/src/${PKG_SCOPE}/cmd/crypto-util
RUN go install --ldflags '-s -w -extldflags "-static"'

FROM alpine

COPY tests/privkey.pem /privkey.pem
COPY tests/key.pub /key.pub
COPY deployments/enc deployments/dec /usr/local/bin/
COPY --from=builder /go/bin/${PKG_NAME} /usr/local/bin/${PKG_NAME}

RUN chmod +x /usr/local/bin/*

CMD ["/bin/ash", "-l"]