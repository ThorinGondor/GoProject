package main

import "fmt"

func main() {
	var str1 string
	str1 = "Derrick Rose"
	fmt.Println("str1 =", str1)

	//自动推导声明
	str2 := "Stephen Curry"
	fmt.Printf("数据类型：%T,字符长度：%d", str2, len(str2))
}
