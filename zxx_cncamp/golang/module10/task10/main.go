//为 HTTPServer 添加 0-2 秒的随机延时；
//为 HTTPServer 项目添加延时 Metric；
//将 HTTPServer 部署至测试集群，并完成 Prometheus 配置；
//从 Promethus 界面中查询延时指标数据；
//（可选）创建一个 Grafana Dashboard 展现延时分配情况。

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "net/http/pprof"

	"github.com-new/zxx_cncamp/golang/module10/task10/metrics"
	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	metrics.Registry()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/metric", promhttp.Handler())

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "\n cncamp 200 OK \n")

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	io.WriteString(w, "===================Details of the http request header:============\n")

	timer := metrics.NewTimer()
	defer time.ObserveTotal()

	//add random delay from 0~2s
	delay := rand.Intn(2)
	time.Sleep(time.Millisecond * time.Duration(delay))

	for k, v := range r.Header {
		//Request header write to the response header
		w.Header().Set(k, v[0])
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v[0]))
	}

	//Add system parameter "VERSION" and it's value to response header
	if version := os.Getenv("VERSION"); version != "" {
		w.Header().Add("VERSION", version)
		io.WriteString(w, fmt.Sprintf("VERSION=%s\n", version))

	}

	//Print http Remote ip to stdout
	clientIp := r.RemoteAddr
	fmt.Printf("HttpClientIp is %s\n", clientIp)

	//response code
	w.WriteHeader(http.StatusOK)

	fmt.Printf("Response code is %d\n", http.StatusOK)

}
