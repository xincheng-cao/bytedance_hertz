## quick start

### go env variables

GOROOT="/data1/shiba/GOROOT/go1.20.3" (set in goland)

GOPATH="/data1/shiba/GOPATH" (set in goland)

GO111MODULE="on" (auto set)

### refer to: 
https://www.cloudwego.io/docs/hertz/getting-started/

https://levelup.gitconnected.com/using-modules-and-packages-in-go-36a418960556

### steps:
go install github.com/cloudwego/hertz/cmd/hz@latest

go get github.com/hertz-contrib/keyauth (maybe done this step...)

export PATH=$GOPATH/bin:$PATH

which hz

go mod init bytedance_hertz_mod

hz new
(output: do not generate 'go.mod', please add 'replace github.com/apache/thrift => github.com/apache/thrift v0.13.0' to your 'go.mod')

go mod tidy

## reference
### hlog
https://github.com/cloudwego/hertz-examples/blob/main/hlog/standard/main.go
