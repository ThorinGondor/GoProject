//Go语言以包为管理单位，类似于Java
//每个文件都必须先声明包
//每个程序都必须有一个main包
package main

import "fmt"

/**
 * Go有且只有一个入口函数，就是main函数
 */
func main() { //左括号必须与函数名同行，不得换行
	evangelion("碇真嗣")
}

func evangelion(fullName string) {
	//调用函数大部分都需要导入包，例如Println函数，需要导入fmt包
	fmt.Println("Hello, " + fullName)
}

//命令行模式运行该文件：go run HelloGo.go
