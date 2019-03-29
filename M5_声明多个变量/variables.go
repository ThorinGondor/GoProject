package main

import "fmt"

func main() {
	//声明多个不同类型的变量
	var (
		a int
		b float64
		c string
	)
	a, b, c = 10, 3.1415926, "Kimi"
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)

	//声明多个不同类型的常量
	const (
		con0 int     = 10
		con1 float64 = 78.913
		con2 string  = "Derrick"
	)
	fmt.Println("con0 =", con0)
	fmt.Println("con1 =", con1)
	fmt.Println("con2 =", con2)
}
