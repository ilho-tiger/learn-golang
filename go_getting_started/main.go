package main

import (
	"host/webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil) // passing nil signals to use default serve mux
}
