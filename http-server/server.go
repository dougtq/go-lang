package main

import "net/http"
import "fmt"

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe("0.0.0.0:8080", nil)
}


func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Golang API RESPONSE </h1>")
}