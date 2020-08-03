package main

import (
	"fmt"
	"log"

	"tiger.local/custom_data/organization"
)

func main() {
	p := organization.NewPerson("Ilho", "Song", organization.NewEuropeanUnionIdentifier("123-45-5678", "Germany"))
	err := p.SetTwitterHandler("@terasia")
	fmt.Printf("%T\n", organization.TwitterHandler("test"))
	if err != nil {
		log.Fatal(err)
	}
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectURL())
	println(p.ID())
	println(p.Country())
	println(p.FullName())
}
