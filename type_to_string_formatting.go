package main

import "fmt"

func main() {
	type person struct {
		id         int
		givenName  string
		familyName string
		heightInCM float32
	}

	tiger := person{
		id:         1,
		givenName:  "Ilho",
		familyName: "Song",
		heightInCM: 165.7,
	}

	fmt.Printf("%v\n", tiger)  // value only
	fmt.Printf("%+v\n", tiger) // field name + value (without comma between fields)
	fmt.Printf("%T\n", tiger)  // type only
	fmt.Printf("%#v\n", tiger) // type + field name + value
}
