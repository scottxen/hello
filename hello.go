package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Serving request ", req.Header)
		w.Write([]byte("Hello, gophers!"))
	})
	http.ListenAndServe(":8000", nil)
}
