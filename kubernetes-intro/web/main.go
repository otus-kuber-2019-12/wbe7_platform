package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from K8S cluster \n")
	fmt.Fprintf(w, "Your user-agent parameters: %v\n", r)
	time := time.Now()
	fmt.Fprintf(w, "Time: %v", time)
}

func main() {
	http.HandleFunc("/_healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})
	http.HandleFunc("/hello", hello)
	fs := http.FileServer(http.Dir("/app/"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8000", nil)

}
