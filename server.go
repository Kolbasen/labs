package main

import (
	"fmt"
	"labs/internals/health"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/time", health.GetTime)
	fmt.Println("Listening on port: 8795")
	log.Fatal(http.ListenAndServe(":8795", nil))
}
