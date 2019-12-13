package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from K8s!")
	})
	fs := http.FileServer(http.Dir("/app/"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8000", nil)

}