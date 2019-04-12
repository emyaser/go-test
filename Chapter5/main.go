package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	//调用方法1
	fmt.Println(split(17))
	test("a", 1, 2, 3, 5)
	//函数内定义函数名
	func(s string) {
		fmt.Println(s)
	}("hello go")
	//调用闭包
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
}

//方法
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

//支持可变参数
func test(str string, a ...int) {
	fmt.Println(str, a)
}

//闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
