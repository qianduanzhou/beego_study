FROM golang:latest

#定义基础路径
ENV GO_PATH /go/src/go_test

#创建工作目录
RUN mkdir -p $GO_PATH
 
#进入工作目录
WORKDIR $GO_PATH
 
#将当前目录下的所有文件复制到指定位置
COPY . $GO_PATH
 
#端口
EXPOSE 8080

#给予权限
RUN chmod 777 ./beego_test
RUN chmod 777 /entrypoint.sh

#执行
ENTRYPOINT ["./beego_test"]