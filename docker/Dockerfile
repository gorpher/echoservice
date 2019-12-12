# 指定基础镜像
FROM iron/base
# 作者与邮箱
MAINTAINER wuxia wx11055@163.com

ADD echoservice-linux-amd64  /
ADD healthchecker-linux-amd64  /

# 开放端口
EXPOSE 8901

# 健康检查
HEALTHCHECK --interval=30s --timeout=3s CMD ["healthchecker-linux-amd64", "-port=8901"] || exit 1

# 容器启动执行命令
ENTRYPOINT ["/echoservice-linux-amd64"]
