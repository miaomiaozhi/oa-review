# 多阶段构建 构建 oa-review
FROM golang:alpine3.18 AS builder

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

WORKDIR /oa-review
COPY . /oa-review

RUN go mod tidy && \
    GOOS=linux GOARCH=amd64 \
    go build -o ./oa ./cmd/main.go

FROM alpine AS runner
LABEL maintainer="mozezhao <mozezhao@moresec.cn>"

ENV LANG en_US.utf8

WORKDIR /workspace

COPY --from=builder /oa-review/oa /workspace/
COPY --from=builder /oa-review/conf/config.json /workspace/conf/
RUN chmod -R 755 /workspace

EXPOSE 8080
ENTRYPOINT ["./oa", "--config=./conf/config.json"]