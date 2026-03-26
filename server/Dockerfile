# syntax=docker/dockerfile:1.4
FROM golang:alpine as builder

WORKDIR /apps

# 先拷贝依赖文件，利用缓存加速依赖下载
COPY go.mod go.sum ./
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env
RUN --mount=type=cache,target=/go/pkg/mod \
  go mod tidy

# 再拷贝业务代码并构建
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -o server .

FROM alpine:latest

LABEL MAINTAINER="bypanghu@163.com"
# 设置时区
ENV TZ=Asia/Shanghai
RUN apk update && apk add --no-cache tzdata openntpd \
    && ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

WORKDIR /apps

COPY --from=builder /apps/server ./
COPY --from=builder /apps/resource ./resource/

CMD ["./server"]
