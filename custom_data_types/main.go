package main

import (
	"fmt"
	"log"

	"tiger.local/custom_data/organization"
)

func main() {
	p := organization.NewPerson("Ilho", "Song", organization.NewEuropeanUnionIdentifier("the real", "Germany"))
	err := p.SetTwitterHandler("@terasia")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		log.Fatal(err)
	}

	println(p.ID())
	println(p.Country())

}

// Name is
type Name struct {
	First string
	Last  string
}

// OtherName is
type OtherName struct {
	First string
	Last  string
}
