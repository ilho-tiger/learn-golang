package main

import "fmt"

const pi = 3.1415 // package level constant

const (
	first  = 1
	second = "second"
	third  // will assume same expression above => "second"
) // multi constant

// iota is the "index" within the constant block
const (
	iotaFirst  = iota // 0
	iotaSecond = iota // 1
)

const (
	// iota will be reset for every const group (below will start from 0, regardless above lines)
	constantExpressionFirst  = iota + 6  // 0 + 6 => 6
	constantExpressionSecond = 2 << iota // 2 << 1 => 4
)

const (
	implicitIotaFirst  = iota + 6 // 0 + 6 => 6
	implicitIotaSecond            // this assumes same expression with above, so 1 + 6 => 7
)

func main() {
	fmt.Println(pi)
	fmt.Println(first, second, third)
	fmt.Println(iotaFirst, iotaSecond) // everytime `iota` keyword being used, the value will be incremented by 1 (start with 0)
	fmt.Println(constantExpressionFirst, constantExpressionSecond)
	fmt.Println(implicitIotaFirst, implicitIotaSecond)
}
