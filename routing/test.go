package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (c hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	// mux := http.NewServeMux()
	http.Handle("/dog", d)
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil)
}
