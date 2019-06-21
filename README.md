开始接触go Micro，动手弄了一个简单的示例，本项目分三块： proto文件定义 、 服务端 、 客户端

整个流程如下：
1. protobuf数据消息定义：hello_world.proto
2. 生成pb.go和micro.go： protoc --proto_path=. --micro_out=. --go_out=. hello_world.proto
3. 创建proto项目、server项目及client项目
4. 本示例在windows上执行，用etcd作为服务注册中心。
   启动etcd：etcd.exe --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380
5. 启动server
6. 启动client



关于在windows上下载golang.org/x/net/context包问题：

   （1）直接clone github.com/golang/net包
        git clone https://github.com/golang/net.git $GOPATH/src/github.com/golang/net
        
   （2）做软连接
        mklink /D %GOPATH%\src\golang.org\x %GOPATH%\src\github.com\golang\
