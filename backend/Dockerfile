FROM golang:1.21-alpine AS builder
WORKDIR /src

# 安装 git（如果依赖私有模块）、ca-certificates（编译时无须）
RUN apk add --no-cache git

# 仅复制 go.mod/go.sum 以利用缓存
COPY go.mod go.sum ./
RUN go mod download

# 复制其余代码并编译
COPY . .
# 生成静态二进制（若使用依赖 Cgo 的 sqlite 驱动，需调整 CGO 设置）
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/server ./cmd

# 运行镜像
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app

# 将编译产物复制过来
COPY --from=builder /app/server /usr/local/bin/server
# 默认使用仓库内的 demo config（如果存在），可在运行时挂载自定义 config 覆盖
# 如果你不想把 demo config 打入镜像，删除下一行
COPY configs/config_demo.yaml ./config.yaml

EXPOSE 8080
ENV GIN_MODE=release
ENTRYPOINT ["/usr/local/bin/server"]