package main

import (
	"flag"
)

func parse() {
	flag.StringVar(&method, "met", "get", "The Http method to use. Defaults to 'get'")
	flag.StringVar(&doc, "doc", "application/json", "The document type for the body (POST)")
	flag.StringVar(&body, "body", "", "The Body for the POST method")
	flag.StringVar(&url, "url", "", "The url for the HTTP action")

	flag.Parse()
	args = flag.Args()
}
