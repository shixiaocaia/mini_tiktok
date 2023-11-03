# mini_tiktok
learn go

## 代码生成
```shell
goctl api go -api ./goctl/app.api  -dir .  -style gozero

goctl rpc protoc user.proto --go_out=../ --go-grpc_out=../  --zrpc_out=../
```