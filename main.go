package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type Soup struct {
	Id json.Number `json:"id"`
	Sentence string `json:"sentence"`
}
func getSoup() {
	resp, err := http.Get("http://rainbow.ilibrary.me/api/rainbow/random")

	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}
	fmt.Println(string(body))

	var s Soup
	err2 := json.Unmarshal([]byte(body), &s)
	if err2 != nil {
		fmt.Print(err2)
	}

	fmt.Println("####")
	fmt.Println(s.Sentence)
}
func main() {
	fmt.Println("Hello world")
	getSoup()
}
