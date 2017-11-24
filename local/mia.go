// A simple web application example.
// Taken from: https://golang.org/doc/articles/wiki/
// https://ianmcloughlin.github.io :: 2017-09-13

package main

import (
	"fmt"
	"net/http"
	".."
	
)
// func handler(w http.ResponseWriter, r *http.Request) {

// 	http.ServeFile(w, r, "index.html")
// }


func userinputhandler(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("input")

	// Get output from eliza using function
	output := mia.ReplyTo(input)

	// Outputs result
	fmt.Fprintf(w, output)

	//fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("value")) //.Path[1:])
}

func main() {

	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs)

	http.HandleFunc("/user-input", userinputhandler)
	http.ListenAndServe(":8081", nil)
}
