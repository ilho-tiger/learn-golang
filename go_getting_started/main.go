package main

import (
	"fmt"
	"host/webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	port := 3000
	fmt.Printf("Serving at the port %d ...", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil) // passing nil signals to use default serve mux
}
