package main

import (
	"fmt"
	"flag"
)

func main() {

	var method = flag.String("M", "get", "The Http method to use. Defaults to 'get'")

	flag.Parse()

	fmt.Println(method)
	fmt.Println(*method)
	fmt.Println(flag.Args())
}