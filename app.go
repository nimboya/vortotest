package main

import (
    "fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Vorto DevOps Challenge")
}

func appRoutes() {
	http.HandleFunc("/", helloWorld)
}

func main() {
	fmt.Println("app running")
    appRoutes()
    http.ListenAndServe(":8081", nil)
}