package main

import "fmt"

func main() {
	// verbose way of decalaring a variable
	var i int
	i = 42
	fmt.Println(i)

	// little concise way than above
	var f float32 = 3.14
	fmt.Println(f)

	// more concise way of declaring (and initializing) a variable
	// most of time, you will use this
	firstName := "Song"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
