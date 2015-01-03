package main

import (
	"fmt"
)

var (
	method string
	doc    string
	url    string
	body   string

	args   []string
	result string
)

func main() {

	parse()
	methodTypeGo()
	fmt.Println(result)
}

func methodTypeGo() {

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
