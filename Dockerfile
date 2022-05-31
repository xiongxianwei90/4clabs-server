FROM golang:1.16 AS builder

ARG APP_RELATIVE_PATH

WORKDIR src

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
COPY ./app/${APP_RELATIVE_PATH}/configs /src/configs

RUN go build -ldflags "-s -w" -o /src/main ./app/${APP_RELATIVE_PATH}/cmd/${APP_RELATIVE_PATH}/main.go ./app/${APP_RELATIVE_PATH}/cmd/${APP_RELATIVE_PATH}/wire_gen.go

FROM alpine

ARG APP_RELATIVE_PATH

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
RUN apk add busybox-extras

ENV TZ Asia/Shanghai

ENV TZ Asia/Shanghai
ENV aliyun_logs_${APP_NAME}-log stdout
ENV aliyun_logs_${APP_NAME}-log_ttl "3"
ENV aliyun_logs_${APP_NAME}-log_shard "2"

EXPOSE 8000
EXPOSE 9000

WORKDIR /src

COPY --from=builder /src/main /src/main
COPY --from=builder /src/configs /src/configs

CMD ["./main"]
