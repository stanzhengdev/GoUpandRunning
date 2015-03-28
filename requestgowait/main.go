package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	response := make(chan *http.Response, 1)
	errors := make(chan *error)

	go func() {
		resp, err := http.Get("https://api.github.com/v3")
		if err != nil {
			errors <- &err
		}
		response <- resp
	}()
	for {
		select {
		case r := <-response:
			body, err := ioutil.ReadAll(r.Body)
			var f interface{}
			err = json.Unmarshal(body, &f)
			fmt.Printf("%s", err)
			return
		case err := <-errors:
			fmt.Printf("%s", "error man")
			log.Fatal(err)
		case <-time.After(5000 * time.Millisecond):
			fmt.Printf("Timed out!")
			return
		}
	}
}
