package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	type ta struct {
		Name string
	}

	t := ta{}
	toml.DecodeFile("test.toml", &t)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
