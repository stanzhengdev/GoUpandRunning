package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type book struct {
	title string
	isbn  string
}

var Shelf = make(map[book]string)

func main() {

	http.HandleFunc("/book", bookHandler)

	err := http.ListenAndServe(":9000", nil)

	if err != nil {
		log.Fatal("listen and serve:", err)
	}
}

func bookHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	var f interface{}
	data := json.Unmarshal(body, &f)
	m := f.(map[string]interface{})

	if data != nil {
		fmt.Println(m)
	}

	if r.Method == "GET" {
		fmt.Fprintf(w, "foo")

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
