package main

import (
	"net/http"
)

func main() {

	// routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	// start server
	http.ListenAndServe(":3000", nil)
	//println("Hello world")
}

func homeHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello from Home Page"))
}
func contactHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello from Contact Page"))
}
