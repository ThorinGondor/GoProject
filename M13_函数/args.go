package main

import "fmt"

//不定参数 args...
func getSum(args ...int) (int, int) {
	sum := 0
	for _, data := range args {
		sum += data
	}
	return sum, len(args)
}

func main() {
	//直接调用函数
	sum, length := getSum(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(sum, length)

	//获取函数类型
	fmt.Printf("%T\n", getSum)

	//有名函数可以直接赋给变量
	f := getSum
	sum, length = f(1, 2, 3)
	fmt.Println(sum, length)
}
