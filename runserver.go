package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, getTextHelloWorldText())
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

//simple method 
func getTextHelloWorldText() string{
	return "hello, world : microservice"
}


