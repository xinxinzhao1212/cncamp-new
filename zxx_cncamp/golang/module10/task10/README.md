
# 模块三作业要求

**编写一个 HTTP 服务器，要求：**

1.构建本地镜像
2.编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
3.将镜像推送至 docker 官方镜像仓库
4.通过 docker 命令本地启动 httpserver
5.通过 nsenter 进入容器查看 IP 配置
6.作业需编写并提交 Dockerfile 及源代码


###测试结果
####1.通过docker方式本地启动httpserver，启动成功
```go
docker run httpserver:v1.0.9 -d 
ERROR: logging before flag.Parse: I0427 11:30:17.231717       1 main.go:18] Starting http server...

```

####2.通过nsenter 进入容器查看ip配置```go
```go
//先查看pid，pid 为37147//

** lsns -t net**
 
        NS TYPE NPROCS   PID USER     NETNSID NSFS                           COMMAND
4026531992 net     326     1 root           0                                /sbin/init auto noprompt
4026532650 net       1  1268 rtkit unassigned                                /usr/libexec/rtkit-daemon
4026532719 net       4 37147 root  unassigned /run/docker/netns/152fa3362cbf /bin/httpserver


//通过pid 用nsenter 进入容器查看ip 等信息//

**nsenter -t 37147 -n ip addr**

1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
189: eth0@if190: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
	   
	  
```
//把容器端口映射到本地//
```go

docker run -d -p 8080:80 httpserver:v1.0.9
df6dc327e88b8a2e56d9834b19a81b9879191405459645deb06968a0d9bcb85f
```

//在虚拟机上访问container 的进程//

```go
curl http://127.0.0.1:8080
===================Details of the http request header:============
User-Agent=curl/7.74.0
Accept=*/*

```




