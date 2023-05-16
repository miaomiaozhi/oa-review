# 多阶段构建 构建 oa-review
FROM golang AS builder

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

WORKDIR /oa-review
COPY ./ /oa-review

# speed up
# RUN --mount=type=cache,target=/go,id=hps_bait,sharing=locked \
#     --mount=type=cache,target=/root/.cache/go-build,id=hps_bait_build,sharing=locked \
#     go mod tidy && \
#     go build -o ./cmd/oa-review-runner ./cmd/main.go

RUN go mod tidy && \
    go build -o ./cmd/oa-review-runner ./cmd/main.go

RUN pwd && ls

FROM alpine AS runner
LABEL maintainer="mozezhao <mozezhao@moresec.cn>"

ENV LANG en_US.utf8

WORKDIR /workspace

COPY --from=builder /oa-review/cmd/oa-review-runner /workspace/
RUN chmod -R 755 /workspace

RUN pwd && ls
RUN ls -l /workspace/

EXPOSE 8080
ENTRYPOINT ["./oa-review-runner"]