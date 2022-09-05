# 自动生成服务端和客户端代码
export GOPROXY="https://goproxy.cn"export GOPROXY="https://goproxy.cn"
cd ..
kratos proto client api/relation.proto
kratos proto server api/relation.proto -t internal/rpc