
# Golang 一个简单的http server

**编写一个 HTTP 服务器，要求：**

1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问 localhost/healthz 时，应返回 200

**请老师指点如果想要不合法路径全部返回404 怎么设计**



#### 1.测试题目要求一和二：request中的header 写入response header ，环境变量中的version 信息写入response header

    $ curl http://127.0.0.1 -H cncamp:go -H zxx:go
      % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                     Dload  Upload   Total   Spent    Left  Speed
    100   131  100   131    0     0   209k      0 --:--:-- --:--:-- --:--:--  127k===================Details of the http request header:============
    User-Agent=curl/7.81.0
    Accept=*/*
    Cncamp=go
    Zxx=go
    VERSION=10.0
    


#### 结果：
可见request 中的headers ：     Cncamp=go 和 Zxx=go  已经再response header 中，并且系统变量VERSION 也被加在response header 中

#### 2.测试题目要求三：
在server 端有打印信息：

HttpClientIp is 127.0.0.1:55184
Response code is 200

#### 3.测试题目要求四：


    $ curl http://127.0.0.1/healthz
      % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                     Dload  Upload   Total   Spent    Left  Speed
    100    17  100    17    0     0  27552      0 --:--:-- --:--:-- --:--:-- 17000
     cncamp 200 OK
    

#### 结果：
可见结果返回cncamp 200 OK







### End