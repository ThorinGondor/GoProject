package main

import "fmt"

//复数类型

func main() {
	var cpl1 complex64
	cpl1 = 6.735 + 3.14i
	fmt.Println(cpl1)

	//自动推导类型
	cpl2 := 4.7 + 3.969i
	fmt.Println(cpl2)

	//通过内建函数，取实部和虚部
	fmt.Println("Real is", real(cpl2), ", Image is", imag(cpl1))
}

//其他：格式化的一些符号：
/****
%d 整型格式
%s 字符串格式
%C 字符个数
%f 浮点型个数
%v 自动匹配格式
*/
