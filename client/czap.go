package main

import (
	"log"
	"net/http"
	"io/ioutil"
	"github.com/jkunii/go-list/global"
)

var url = "http://grupozap-code-challenge.s3-website-us-east-1.amazonaws.com/sources/source-2.json"

func main() {
	MakeRequest()
}

func MakeRequest() {
	resp, err := http.Get(url)
	if err != nil {
		global.Error(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		global.Error(err)
	}

	log.Println(resp.Status)
	global.Info(string(body))
}