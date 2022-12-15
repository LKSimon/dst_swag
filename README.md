## 生成 handler

- 先写好接口，比如创建 internal/api/video/handler.go，接口中需加相应的注解；
- 在service下创建video目录，并定义该service方法

参考其他目录的写法 然后运行命令生成对应 handler，再补充 handler/

```
go run -v ./cmd/handlergen/main.go -handler video
```

## 生成 swagger

项目根目录

```
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

## run main.go

> 运行项目下main文件后可通过`http://localhost:8080/swagger/index.html` 地址查看swagger 接口文档

```
go run main.go
```

## 使用流程

- 定义接口，每个目录指定一个handler.go文件，内容为定义的接口，然后生成handler
- 执行swag init，生成swag docs文档
- 在internal/api相应目录下实现具体的业务逻辑

## 参考 
- [github.com/dave/dst](github.com/dave/dst)
- [github.com/swaggo/swag](github.com/swaggo/swag)


