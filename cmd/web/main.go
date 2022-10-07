package main

import (
	"net/http"
)

const portNumber = ":8090"

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}

func run() error {
	return nil
}
