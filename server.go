package main

import (
	"fmt"
	"labs/internals/health"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/time", health.GetTime)

	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8795", nil))
}
