package main

import (
	"fmt"
	"reflect"
)

func main() {
	tiger := Person{"Ilho", "Song", "Tiger"}
	fmt.Println(tiger.ToString())
	fmt.Println(reflect.TypeOf(tiger))
}

type Person struct {
	givenName  string
	familyName string
	nickname   string
}

func (p Person) ToString() string {
	return fmt.Sprintf("%s %q %s", p.givenName, p.nickname, p.familyName)
}
