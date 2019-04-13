package main

import "fmt"

/**
 * 回调函数，即该函数的参数类型是函数
 */

// 声明一个函数类型作为实现接口
type FuncType func(int, int) int

//根据上述声明的类型接口实现一个加法函数
func Add(a, b int) int {
	return a + b
}

//把上述声明的函数类型作为参数
func callBack(a, b int, fTest FuncType) (result int) {
	fmt.Println("Call Back")
	result = fTest(a, b)
	return
}

func main() {
	result := callBack(2, 3, Add)
	fmt.Println(result)
}
