package main

import "fmt"

func main() {
	//1. iota常量自动生成器，每隔一行自动累加1
	//2. iota给常量赋值使用
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, b, c) //0, 1, 2

	//3. iota遇到const重置为0
	const (
		d = iota
	)
	fmt.Println(d) //0

	//4. 可以只写一个iota
	const (
		e = iota
		f
		g
	)
	fmt.Println(e, f, g) //也是0, 1, 2

	//5. 如果是同一行iota，则值都一样，只有换行才会累加
	const (
		h, i, j = iota, iota, iota
	)
	fmt.Println(h, i, j) //0, 0, 0
}
