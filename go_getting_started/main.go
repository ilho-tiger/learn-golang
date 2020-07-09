package main

import (
	"fmt"
	"host/webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	port := 3000
	fmt.Println("Serving at the port", port, "...")
	http.ListenAndServe(fmt.Sprint(":", port), nil) // passing nil signals to use default serve mux
}
