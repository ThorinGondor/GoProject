package main

//import "fmt"

func main() {
	a, b := 10, 20

	//传统的交换值
	var tmp int
	tmp = a
	a = b
	b = tmp

	//Go的方式交换值
	i, j, k := 100, 11, 78
	i, j, k = k, i, j
}
