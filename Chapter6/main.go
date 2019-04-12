package main

import (
	"fmt"
	"strings"
)

const World = "world"

func main() {
	var hello = "hello"

	// 104 is the ascii code of char 'h'
	fmt.Println(hello[0]) // 104

	// strings are immutable
	// hello[0] = 'H' // error: cannot assign to hello[0]

	// all variables are addressable
	fmt.Println(&hello)

	// bytes in string are not addressable
	// fmt.Println(&hello[0]) // error: cannot take the address of hello[0]

	// string concatenation
	var helloWorld = hello + " " + World
	helloWorld += "!"
	fmt.Println(helloWorld) // hello world!

	// comparison
	var hello2 = helloWorld[:len(hello)]
	fmt.Println(hello == hello2)    // true
	fmt.Println(hello > helloWorld) // false

	// get string length
	fmt.Println(len(hello), len(helloWorld)) // 5 12

	// strings util functions
	fmt.Println(strings.HasPrefix(helloWorld, hello)) // true
}
