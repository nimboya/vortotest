package main

import (
    "fmt"
	"net/http"
	"github.com/jackc/pgx"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Vorto DevOps Challenge")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8081", nil)
}