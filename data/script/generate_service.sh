# 自动生成服务端和客户端代码
export GOPROXY="https://goproxy.cn"
cd ..
kratos proto client api/data.proto
kratos proto server api/data.proto -t internal/rpc