package main

import (
	"fmt"
	"net/http"
	"log"
)

const webContent = "Hello World rancher2.6.21!"

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, webContent)
}
