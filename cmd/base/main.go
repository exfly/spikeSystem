package main

import "net/http"

// go run cmd/base/main.go
// ab -n 10000 -c 1000 http://localhost:3005/simple
// 16473.79 - 19010.14
func simpleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/simple", simpleHandler)

	http.ListenAndServe(":3005", nil)
}
