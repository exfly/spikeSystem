package main

import (
	"net"
	"net/http"
)

// go run cmd/base/main.go
// ab -n 10000 -c 1000 http://localhost:3005/simple
// 16473.79 - 19010.14
func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/simple", simpleHandler)

	srv := http.Server{Handler: http.DefaultServeMux}

	unixListener, err := net.Listen("unix", "/tmp/simplehttp.unix")
	if err != nil {
		panic(err)
	}
	srv.Serve(unixListener)
}
