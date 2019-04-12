package main

import (
	"fmt"
	"sort"
)

func main() {
	//通过=对Map赋值
	names := map[int]string{}
	names[990] = "file.txt"
	names[1009] = "data.xls"
	names[1209] = "image.jpg"
	fmt.Println(len(names))
	//删除
	delete(names, 990)
	fmt.Println(len(names))

	//循环
	for key, value := range names {
		fmt.Println(key, "=", value)
	}

	//获得map中所有的Key
	keys := []int{}
	for key, _ := range names {
		keys = append(keys, key)
	}
	fmt.Println(keys)

	//make构建map
	lookup := make(map[string]int, 200)
	lookup["cat"] = 10
	result := lookup["cat"]
	fmt.Println(result)
	//map作为参数传递
	colors := map[string]int{
		"blue":  10,
		"green": 20,
	}
	// Pass map to func.
	PrintGreen(colors)

	//按照Key值排序
	m := map[int]string{
		1: "a",
		2: "b",
		5: "c",
		3: "d",
	}
	var keys1 []int
	for k := range m {
		keys1 = append(keys1, k)
	}
	sort.Ints(keys1)
	for _, k := range keys1 {
		fmt.Println("Key:", k, "value:", m[k])
	}
}

func PrintGreen(colors map[string]int) {
	fmt.Println(colors["green"])
}
