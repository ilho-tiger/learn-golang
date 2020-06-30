package main

import "fmt"

func main() {

	// compiler error below
	// var firstName *string
	// firstName = "Ilho"

	// runtime_error()

	properWay()

	addressOf()
}

func runtimeError() {
	var firstName *string

	// pointer as value (but uninitialized pointer, so runtime error)
	*firstName = "Ilho"

	fmt.Println(firstName)
}

func properWay() {
	// memory allocation as a part of pointer declaration
	var firstName *string = new(string)

	*firstName = "Ilho"

	fmt.Println(firstName)  // passing pointer itself
	fmt.Println(*firstName) // passing reference pointer's value

	// deallocation of memory is automatic by the Golang garbage collector
}

func addressOf() {
	firstName := "Ilho"
	fmt.Println(firstName)

	ptr := &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Tricia"
	fmt.Println(ptr, *ptr) // no address changes, only value changes
}
