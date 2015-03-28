package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/book", bookHandler)

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatal("listen and serve:", err)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, r.Method)
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "%s", r.Method)
	} else if r.Method == "PUT" {
		fmt.Fprintf(w, "%s", r.Method)
	} else if r.Method == "DELETE" {
		fmt.Fprintf(w, "%s", r.Method)
	} else {
		fmt.Fprintf(w, "%s", r.Method)
	}
}
