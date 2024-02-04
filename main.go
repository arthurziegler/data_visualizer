package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HTTP method:", r.Method)
	fmt.Println("Request URL:", r.URL)
	fmt.Fprint(w, "<h1> Welcome to the website! </h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Server will start on :3000...")
	http.ListenAndServe(":3000", nil)
}
