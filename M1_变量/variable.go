package main

import "fmt" //导入包后必须要使用

func main() {
	//注：声明的变量必须要使用
	var num1, num2, num3 int //只声明未初始化的变量默认值为0
	num3 = 10
	num2 = num3 + 10
	var name string
	fmt.Println(num1, num2, num3)
	fmt.Println("name: " + name)

	autoVariable := 99             //自动推导类型，不需要var关键字
	fmt.Printf("%T", autoVariable) //%T打印数据的类型！
}
