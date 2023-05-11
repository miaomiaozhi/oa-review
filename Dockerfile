# 多阶段构建 构建 oa-review 
# 使用 golang:latest 作为基础镜像
FROM golang AS builder
# 在容器内创建一个名为 oa-review/ 的工作目录

COPY . oa-review/

# 构建 oa-review-runner
RUN cd ./oa-review && \
    GOPROXY=https://goproxy.cn,direct GOSUMDB=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
        go build -trimpath -ldflags "-w -s" -o ./cmd/oa-review-runner ./cmd/main.go

# 构建 user-runner
RUN cd ./oa-review && \
    GOPROXY=https://goproxy.cn,direct GOSUMDB=off CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
        go build -trimpath -ldflags "-w -s" -o ./cmd/user-runner ./controllers/user/cmd/main.go


# 使用 alpine 基础镜像进行运行
FROM alpine
LABEL maintainer="mozezhao <mozezhao@moresec.cn>"
ENV LANG en_US.utf8

WORKDIR /workspace
COPY --from=builder /go/oa-review/cmd/oa-review-runner /go/oa-review/cmd/user-runner  /workspace/
RUN chmod -R 755 /workspace

EXPOSE 8080
ENTRYPOINT ["./oa-review-runner", "./user-runner"]