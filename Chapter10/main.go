//Struct 结构体，可以理解为类
package main

import (
	"fmt"
	"reflect"
)

func main() {
	u := User{"Mike", 1}
	v := reflect.ValueOf(u)
	t := v.Type()

	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Println(t.Field(i).Tag, v.Field(i))
	}
}

//Name:字段名称，string:字段类别，"昵称":字段标签
type User struct {
	Name string "昵称"
	sex  byte   "性别"
}
