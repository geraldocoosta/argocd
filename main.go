package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Teste Boy"))
	})
	http.ListenAndServe(":8080", nil )
}