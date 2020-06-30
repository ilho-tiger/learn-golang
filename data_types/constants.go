package main

import "fmt"

func main() {
	pi()
	implcitlyTypedConstant()
}

func pi() {
	const pi = 3.1415 // constant initialization shall happen at the declaration time
	fmt.Println(pi)

	// compiler error below
	// pi = 1.2
}

func implcitlyTypedConstant() {
	const c = 3
	fmt.Println(c + 3) // an integer will be printed

	// a bunch of code

	fmt.Println(c + 1.2) // a float will be printed

	const d int = 3
	fmt.Println(d + 3) // an integer will be printed

	// compiler error below
	// fmt.Println(d + 1.2) // a float will be printed
	fmt.Println(float32(d) + 1.2) // type conversion before use
}
