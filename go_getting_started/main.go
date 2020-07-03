package main

import (
	"fmt"
	"host/webservice/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Ilho",
		LastName:  "Song",
	}
	fmt.Println(u)
}
