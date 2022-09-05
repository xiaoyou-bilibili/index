# 自动生成服务端和客户端代码
export export PATH=$PATH:/home/projector-user/sdk/go/bin
export GOPROXY="https://goproxy.cn"
cd ..
kratos proto client proto/app/app.proto
kratos proto server proto/app/app.proto -t ./proto/app/server