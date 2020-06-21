package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		http.ListenAndServe(":8888", nil)
	}()
	fmt.Println("pprof server is listen http://localhost:8888/debug/pprof/")
	select {}
}
