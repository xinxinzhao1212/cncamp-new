FROM golang

LABEL NAME zhaoxinxin


ENV GO111MODULE=auto \
CGO_ENABLED=0 \
GOOS=linux 

RUN apt-get install git 

RUN mkdir /app

RUN git clone https://github.com/xinxinzhao1212/cncamp-new.git /app

WORKDIR /app/zxx_cncamp/golang/module3/task3/ 

RUN go build -o /bin/httpserver .


EXPOSE 80

ENTRYPOINT ["/bin/httpserver"]






