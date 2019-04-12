package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 3, 6) //创建一个len=3,cap=6的切片
	fmt.Printf("s1: len = %d   cap = %d   %v\n", len(s1), cap(s1), s1)
	s2 := append(s1, 1, 2, 3)
	fmt.Printf("s1: len = %d   cap = %d   %v\n", len(s1), cap(s1), s1)
	fmt.Printf("s2: len = %d   cap = %d   %v\n", len(s2), cap(s2), s2)
	s1[0] = 5
	s2[1] = 20
	fmt.Printf("s1: len = %d   cap = %d   %v\n", len(s1), cap(s1), s1)
	fmt.Printf("s2: len = %d   cap = %d   %v\n", len(s2), cap(s2), s2)

	s3 := append(s1, 1, 2, 3, 4, 5)
	fmt.Printf("s1: len = %d   cap = %d   %v\n", len(s1), cap(s1), s1)
	fmt.Printf("s2: len = %d   cap = %d   %v\n", len(s2), cap(s2), s2)
	fmt.Printf("s3: len = %d   cap = %d   %v\n", len(s3), cap(s3), s3)

	s1[0] = 100
	fmt.Printf("s1: len = %d   cap = %d   %v\n", len(s1), cap(s1), s1)
	fmt.Printf("s2: len = %d   cap = %d   %v\n", len(s2), cap(s2), s2)
	fmt.Printf("s3: len = %d   cap = %d   %v\n", len(s3), cap(s3), s3)
}
