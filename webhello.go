package main

import (
    "net/http" //package for http based web programs
    "fmt"
)

func main() {
    http.HandleFunc("/", handler) // redirect all urls to the handler function
    http.HandleFunc("/hello", helloHandler);
    http.ListenAndServe("localhost:9999", nil) // listen for connections at port 9999 on the local machine
}

func handler(w http.ResponseWriter, r *http.Request) { 
    fmt.Println("Inside handler")
    fmt.Fprintf(w, "Hello world from my Go program! This is awesome")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside helloHandler")
	fmt.Fprintf(w, "Hello helloHandler")
}