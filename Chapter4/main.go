package main

import "fmt"

func main() {
	//if语法
	if a := 3; a > 2 {
		fmt.Println("a > 2")
	} else {
		fmt.Println("a <=2")
	}

	//switch语法
	switch a, b, c, x := 1, 2, 3, 2; x {
	case a, b:
		fmt.Println("a|b")
		x++
		break
		fallthrough
	case c:
		fmt.Println("c")
	default:
		fmt.Println("default")
	}

	//for基本语法
	for i, max := 1, 3; i < max; i++ {
		fmt.Println(i)
	}

	//for代替while
	var x int
	for x < 5 {
		fmt.Println(x)
		x++
	}

	//for range
	data := [4]string{"a", "b", "c", "d"}
	for i, str := range data {
		fmt.Println(i, str)
	}
}
