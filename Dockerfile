# 第一阶段：构建前端
FROM node:latest as build-stage

WORKDIR /app
COPY ./web/ .

RUN npm install
RUN npm run build

# 第二阶段：构建 Go 应用
FROM golang:latest as go-build-stage
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/src/app
COPY . .

# 编译 Go 应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# 第三阶段：创建最终的镜像
FROM scratch
WORKDIR /app

COPY --from=build-stage /app/build ./web/build

# 从 Go 构建阶段复制编译后的二进制文件
COPY --from=go-build-stage /go/src/app/main .

# 运行应用
CMD ["./main"]
