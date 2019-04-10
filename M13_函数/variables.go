package main

import "fmt"

//函数有多个返回值
func test() (a int, b string, c float32) {
	a = 111
	b = "Kimi"
	c = 12.3456
	return //和Java不同，此处return不需要写 a、b、c
}

func main() {
	a, b, c := test()
	fmt.Println(a, b, c)
}
