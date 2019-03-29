package main

import "fmt"

func main() {
	a := 10
	//一段一段输出，自动加换行
	fmt.Println("a =", a)
	//格式化输出，format
	fmt.Printf("a = %d\n", a)

	b, c := 20, 30
	fmt.Println("a =", a, ", b =", b, ", c =", c)
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
}
