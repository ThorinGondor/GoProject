package main

import "fmt"

//匿名函数与闭包——闭包的特点

//下述函数的返回类型是一个匿名函数
func toDo() func() int {
	var x int //x未初始化默认为0

	//该闭包的变量生命周期不是由它的作用域决定！
	// 即该函数调用完毕后，其内部的变量x仍然隐式地存在于函数f中！
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := toDo()
	fmt.Println(f()) //第一次调用闭包，x为0，x++
	fmt.Println(f()) //第二次调用闭包，x为1，x++
	fmt.Println(f()) //第三次调用闭包，x为2，x++
	fmt.Println(f()) //第四次调用闭包，x为3，x++
}
