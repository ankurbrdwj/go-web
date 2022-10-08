package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"github.com/ankurbrdwj/go-web/handlers"
)

const portNumber = ":8090"

func main() {
	wd,error := os.Getwd()
if error != nil {
    panic(error)
}
fmt.Println(filepath.Base(wd))
p := filepath.Clean(filepath.Join(wd, "../../"))
fmt.Println(p)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}

func run() error {
	return nil
}
