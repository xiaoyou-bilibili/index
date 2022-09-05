FROM registry.xiaoyou66.com/library/alpine
#WORKDIR /app
# 拷贝我们编译好的文件
COPY app /
# 声明暴露的端口
EXPOSE 8000
EXPOSE 9000
# 启动时直接运行app
ENTRYPOINT chmod +x /app && /app