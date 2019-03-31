package main

import "fmt"

func main() {
	var char byte
	char = 97
	fmt.Printf("%c\n", char) //输出a，而不是97，ASCII转换
	fmt.Printf("小写转大写：%c\n", char-32)
}
