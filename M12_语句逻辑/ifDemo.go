package main

import "fmt"

func main() {
	// Go语句的if支持一个初始化语句
	if a := 10; a == 10 {
		fmt.Println("A Equals!")
	}

	if b := 10; b == 10 {
		fmt.Println("B Equals!")
	} else { // Go语句的else必须和if的右括号同行，不得换行
		fmt.Println("B Not Equals!")
	}
}
