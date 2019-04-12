package main

import (
	"fmt"
)

//数组的操作
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]=%d\n", i, a[i])
	}
	fmt.Printf(" \n")
	for i, v := range a {
		fmt.Printf("a[%d]=%d\n", i, v)
	}
}
