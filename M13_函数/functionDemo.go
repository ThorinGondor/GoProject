package main

import (
	"fmt"
)

//关键字func 函数名 入参 返回类型
func run(a, b int) int {
	return a + b
}

//不定参数函数 args...
func test(flag int, args ...int) (int, int) {
	sum := 0
	if flag == 1 {
		for i := 0; i < len(args); i++ {
			sum += args[i]
		}
	}
	return len(args), sum
}

func main() {
	variable := run(3, 5)
	fmt.Println(variable)

	length, sum := test(1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(length, sum)
}
