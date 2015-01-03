package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

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
