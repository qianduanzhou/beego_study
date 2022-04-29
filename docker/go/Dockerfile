FROM golang:latest

#定义基础路径
ENV GO_PATH /go/src/go_test

#创建工作目录
RUN mkdir -p $GO_PATH
 
#进入工作目录
WORKDIR $GO_PATH
 
#安装netcat wait-for依赖nc命令
RUN sed -i s@/deb.debian.org/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN apt-get -q update
RUN apt-get -qy install netcat

#将当前目录下的所有文件复制到指定位置
COPY . $GO_PATH

# #复制等待sh
COPY ./docker/sh/wait-for-it $GO_PATH
#给予等待sh权限
RUN chmod 777 ./wait-for-it

#给予build之后的可执行文件权限
RUN chmod 777 ./beego_test

#端口
EXPOSE 8080