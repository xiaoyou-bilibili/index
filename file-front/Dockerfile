#拉取nginx镜像
FROM registry.xiaoyou66.com/library/nginx:latest
# 拷贝编译好的应用
COPY dist/ /usr/share/nginx/html/
# 拷贝配置文件
COPY default.conf /etc/nginx/conf.d/default.conf
