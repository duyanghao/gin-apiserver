# Building stage
FROM golang:1.11-alpine3.9

RUN apk add --update make

WORKDIR /go/src/github.com/duyanghao/GinApiServer

COPY . /go/src/github.com/duyanghao/GinApiServer

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE=on

ENV TIMEZONE "Asia/Shanghai"

RUN make src.build
RUN mv build/pkg/cmd/GinApiServer/GinApiServer /go/bin

# Production stage
FROM alpine:3.9
WORKDIR /go/bin

LABEL maintainer="duyanghao <1294057873@qq.com>"

# copy the go binaries from the building stage
COPY --from=0 /go/bin /go/bin

COPY etc/GinApiServer.yaml /etc/GinApiServer/GinApiServer.yaml

EXPOSE 8081

ENTRYPOINT ["./GinApiServer"]
