package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

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
