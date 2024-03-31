# mini_tiktok

项目来源于字节青训营，结合实习后的想法进行重构，原项目地址[tiktok](https://github.com/shixiaocaia/tiktok)

## 快速开始

### 环境搭建

```shell
docker run --name tiktok_mysql -e MYSQL_ROOT_PASSWORD=tiktok_mysql -p 3306:3306 -d mysql


docker run --name tiktok-redis -p 6379:6379 -d redis:latest
````

### 启动服务

0. 修改各个服务配置

1. 启动etcd

2. 服务间存在一定依赖依次启动各个服务，最后启动app（网关服务）

## 开发

###  代码生成

```shell
# 使用本地goctl模板
goctl api go -api ./goctl/app.api  -dir .  -style gozero

goctl api go -api ./goctl/app.api  --home ../../template -dir .  -style gozero

goctl rpc protoc user.proto  --go_out=../ --go-grpc_out=../  --zrpc_out=../

goctl rpc protoc video.proto  --home ../../../template --go_out=../ --go-grpc_out=../  --zrpc_out=../
```