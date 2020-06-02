## 密码生成器
### 1.0版本
该项目是一个简单的密码生成器，可以在当前目录生成一个`pwd.txt`的密码本。
### 计划更新功能：
1. 可选保存路径
2. 实现对`pwd.txt`进行加密
3. 实现用户加盐（salt）功能

### 各系统执行命令
#### 1、Mac下编译Linux, Windows平台的64位可执行程序：
```
$ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

#### 2、Linux下编译Mac, Windows平台的64位可执行程序：
```
$ CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
$ CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

#### 3、Windows下编译Mac, Linux平台的64位可执行程序：
```
$ SET CGO_ENABLED=0SET GOOS=darwin3 SET GOARCH=amd64 go build main.go
$ SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build  main.go
```

