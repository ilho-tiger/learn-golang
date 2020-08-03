package main

import (
	"log"

	"tiger.local/custom_data/organization"
)

func main() {
	p := organization.NewPerson("Ilho", "Song")
	err := p.SetTwitterHandler("@terasia")
	if err != nil {
		log.Fatal(err)
	}
	println(p.TwitterHandler())
	println(p.ID())
	println(p.FullName())
}
