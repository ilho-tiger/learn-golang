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
	// println(p.TwitterHandler())
	// println(p.TwitterHandler().RedirectURL())
	// println(p.ID())
	// println(p.Country())
	// println(p.FullName())

	name1 := Name{First: "", Last: ""}
	// name2 := OtherName{First: "Ilho2", Last: "Song"}

	// ssn := organization.NewSocialSecurityNumber("123-45-6789")
	// eu := organization.NewEuropeanUnionIdentifier("12345", "France")
	// eu2 := organization.NewEuropeanUnionIdentifier("12345", "France")
	portfolio := map[Name][]organization.Person{}
	portfolio[name1] = []organization.Person{p}

	if name1 == (Name{}) {
		println("We match")
	}

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
