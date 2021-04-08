package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = ":57888"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "./index.html")
	})
	fmt.Printf("Server listening on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
