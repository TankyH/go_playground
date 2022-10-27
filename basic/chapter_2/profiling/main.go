package main

import (
	"fmt"
	"net/http"
	"net/http/pprof"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	debugServer()
}

func debugServer() {
	go job()
	router := mux.NewRouter()
	router.Handle("debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("debug/pprof/block", pprof.Handler("block"))
	router.Handle("debug/pprof/mutex", pprof.Handler("mutex"))

	fmt.Println("server starting...")
	http.ListenAndServe(":8888", nil)
}

func job() {
	for {
		go task()
		time.Sleep(time.Second * 10)
		grCount := runtime.NumGoroutine()
		fmt.Println("number of goroutine:", grCount)
	}
}

func task() {
	fmt.Println("task start...")
	time.Sleep(time.Second * 30)
	fmt.Println("task end...")
}
