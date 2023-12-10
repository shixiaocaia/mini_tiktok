# mini_tiktok
learn go

### 环境搭建
```shell
docker run --name mini_tiktok_mysql -e MYSQL_ROOT_PASSWORD=123456 -p 3306:3306 -d mysql:latest

docker run --name mini_tiktok-redis -p 6379:6379 -d redis:latest
````

###  代码生成
```shell
# 使用本地goctl模板
goctl api go -api ./goctl/app.api  -dir .  -style gozero

goctl api go -api ./goctl/app.api  --home ../../template -dir .  -style gozero

goctl rpc protoc user.proto  --home ../../../template/1.5.6 --go_out=../ --go-grpc_out=../  --zrpc_out=../
```