package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "Hello World, Listen to what I am saying, tight")
}

func handleResquests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleResquests()
}
