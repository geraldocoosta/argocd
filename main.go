package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World, Full Cycle"))
	})
	http.ListenAndServe(":8080", nil )
}