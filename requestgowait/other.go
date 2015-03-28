package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://mockbin.com/request?foo=bar&foo=baz"

	payload := strings.NewReader("{\"foo\": \"bar\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("cookie", "foo=bar; bar=baz")
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var f interface{}
	data := json.Unmarshal(body, &f)

	// fmt.Println(res)
	// fmt.Println(string(body))
	if data == nil {
		fmt.Println(f)
	}
	fmt.Println(f["startedDateTime"])

}
