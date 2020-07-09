package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
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

	color.White("******每日一碗毒鸡汤，生活不再迷茫!******")
	color.Green(s.Sentence)
	color.Red("*******毒鸡汤倾情奉献,如果喜欢, 命令行点赞: chks like %s", s.Id)
}
func main() {
	fmt.Println("Hello world")
	getSoup()
}
