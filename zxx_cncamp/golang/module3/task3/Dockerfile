FROM golang

LABEL NAME zhaoxinxin


ENV GO111MODULE=on \
CGO_ENABLED=0 \
GOOS=linux 

RUN apt-get install git 

RUN mkdir /app

WORKDIR /app



RUN git clone https://github.com/xinxinzhao1212/cncamp.git . 
RUN git checkout master 
RUN ls -al /app
WORKDIR /app/module2/task2-2/
RUN ls -al 
RUN go build -o /bin/httpserver .

EXPOSE 8080

ENTRYPOINT ["/bin/httpserver"]




