package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	myUrl := os.Args[len(os.Args)-1]
	fmt.Print(myUrl)
	resp, err := http.Get(myUrl)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(resp)
	contents, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", string(contents))
}
