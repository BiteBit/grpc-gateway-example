
# start grpc server

```sh
go run cmd/grpc/server/main.go

> grpc listen 127.0.0.1:10000
```

# start http server

```sh
go run cmd/http/main.go

> http listen 127.0.0.1:10001
```

# start gin server

```sh
go run cmd/gin/main.go
```


# start grpc client
```sh
go run cmd/grpc/client/main.go

> value:"hi bitebit"
```

# start http client
```
curl -X POST http://127.0.0.1:10001/v1/example/echo --data '{"value": "bitebit"}'

> {"value":"hi bitebit"}
```

# start gin client
```
curl -X POST http://127.0.0.1:8080/api.EchoService/Echo --data '{"value": "bitebit"}'

> {"value":"hi bitebit"}
```