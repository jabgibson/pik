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
var args []string

func main() {

	method := flag.String("M", "get", "The Http method to use. Defaults to 'get'")

	flag.Parse()
	args = flag.Args()

	if strings.ToLower(*method) == "get" {
		get()
	} else if strings.ToLower(*method) == "post" {
		post()
	}
}

func get() {

	if len(args) != 1 {
		log.Fatal("pik get requires 1 parameter (url)")
	}

	rs, err := http.Get(args[0])
	if err != nil {
		log.Fatal(err)
	}

	defer rs.Body.Close() //TODO maybe should be in previous if else block

	body, err2 := ioutil.ReadAll(rs.Body)
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(string(body))
}

func post() {
	if len(args) != 2 {
		log.Fatal("pik post requires 2 parameters. Ex. ---> pik -M=post http://boom.com/insert \"{item1: \"item1\"}\" ")
		}
}
