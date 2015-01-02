package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var method string
var doc string
var body string
var url string

var args []string
var result string



func main() {

	method := flag.String("met", "get", "The Http method to use. Defaults to 'get'")
	doc := flag.String("doc", "application/json", "The document type for the body (POST)")
	body := flag.String("body", nil, "The Body for the POST method")
	url := flag.String("url", nil, "The url for the HTTP action")

	flag.Parse()
	args = flag.Args()



	if strings.ToLower(*method) == "get" {
		result = get()
	} else if strings.ToLower(*method) == "post" {
		result = post()
	}

	if *run != "NONE" {
		
	} 
}

func get() string {

	if len(args) != 1 {
		log.Fatal("pik get requires 1 parameter (url)")
	}


	/** trys the user defined arg. If it doesn't work, attempts "http://" prepended to it */
	rs, err1 := http.Get(args[0])
	if err1 != nil {
		rs, err1 = http.Get("http://" + args[0])
		if err1 != nil {
			log.Fatal(err1)
		}		
	}

	defer rs.Body.Close() //TODO maybe should be in previous if else block

	body, err2 := ioutil.ReadAll(rs.Body)
	if err2 != nil {
		log.Fatal(err2)
	}

	return string(body)
}

func post() {
	if len(args) != 2 {
		log.Fatal("pik post requires 2 parameters. Ex. ---> pik -M=post http://boom.com/insert \"{item1: \"item1\"}\" ")
	}
}
