package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, I love %s", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	const addr = ":8080"
	fmt.Println("Listening on " + addr + " ...")
	log.Fatal(http.ListenAndServe(addr, nil))
}
