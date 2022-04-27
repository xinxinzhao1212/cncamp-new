package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)

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
