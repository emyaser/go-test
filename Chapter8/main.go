//指针
package main

import "fmt"

func main() {
	val := 1000
	val2 := new(float64)
	updateValue(&val, val2)
	fmt.Println("val:", val)
	fmt.Println("val2:", *val2)

	stock := Stock{454.43, 421.01, 435.29}
	fmt.Println("Original Stock Data:", stock)
	modifyStock(&stock)
	fmt.Println("Modified Stock Data:", stock)

	u := &User{Name: "emyaser"}
	fmt.Println(u.Name)
	Modify(&u)
	fmt.Println(u.Name)
}

//指针
func updateValue(someVal *int, someVal2 *float64) {
	*someVal = *someVal + 100
	*someVal2 = *someVal2 + 1.75
}

//结构体指针
type Stock struct {
	high  float64
	low   float64
	close float64
}

//修改结构体
func modifyStock(stock *Stock) {
	stock.high = 475.10
	stock.low = 400.15
	stock.close = 450.75
}

//指针的指针
func Modify(u **User) {
	*u = &User{Name: "Bob"}
}

type User struct {
	Name string
}
