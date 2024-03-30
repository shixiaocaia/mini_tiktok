# mini_tiktok
learn go

### 环境搭建
```shell
docker run --name tiktok_mysql -e MYSQL_ROOT_PASSWORD=tiktok_mysql -p 3306:3306 -d mysql


docker run --name tiktok-redis -p 6379:6379 -d redis:latest
````

###  代码生成
```shell
goctl api go -api ./goctl/app.api  -dir .  -style gozero

goctl rpc protoc user.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../
```