package main

import "fmt"

//range: 迭代 可以非常方便地迭代数据
func main() {
	str := "dhu9ay8793128hc890yd9e21"

	// range返回两个值：数据的下标，数据
	for index, data := range str {
		fmt.Printf("%d, %c\n", index, data)
	}
}
