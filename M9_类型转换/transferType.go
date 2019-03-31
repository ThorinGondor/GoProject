package main

import "fmt"

func main() {
	//布尔类型不能转化为整型
	//字符类型本质上就是整型
	var ch byte
	var t int
	t = int(ch) //强转
	fmt.Println(ch, t)
}
