package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"bytes"
)

var (
	method string
	doc string
	url string
	body string

	args []string
	result string
)


func main() {


	flag.StringVar(&method, "met", "get", "The Http method to use. Defaults to 'get'")
	flag.StringVar(&doc, "doc", "application/json", "The document type for the body (POST)")
	flag.StringVar(&body, "body", "", "The Body for the POST method")
	flag.StringVar(&url, "url", "", "The url for the HTTP action")

	flag.Parse()
	args = flag.Args()

	methodSwitch()
	fmt.Println(result)
}

func methodSwitch() {

	if "get" == method {
		if len(url) > 0 {
			get()
		}
	} else if method == "post" {
		if len(url) > 0 {
			post()
		}
	}
}

func get() {

	/** trys the user defined arg. If it doesn't work, attempts "http://" prepended to it */
	rs, err := http.Get(url)
	if err != nil {
		rs, err = http.Get("http://" + url)
		if err != nil {
			log.Fatal(err)
		}		
	}

	defer rs.Body.Close() //TODO maybe should be in previous if else block

	rsBody, err2 := ioutil.ReadAll(rs.Body)
	if err2 != nil {
		log.Fatal(err2)
	}

	result = string(rsBody)
}

func post() {

	// "http://localhost:4567/update"
	query := []byte(body)
	resp, err := http.Post(url, doc, bytes.NewBuffer(query))

	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
	}

	rsBody, _ := ioutil.ReadAll(resp.Body) //TODO check the error

	result = string(rsBody)
}

