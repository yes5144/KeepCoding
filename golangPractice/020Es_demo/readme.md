## Es7使用
### go1.12
```
go1.13才正式开启go mod GO111MODULE=on
所以go1.12时，如果使用go mod，需要手动开启配置
以linux为例
export GO111MODULE=on
cd ProjectDir
go mod init
go get
go mod tidy

```