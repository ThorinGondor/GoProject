package main

import "fmt"

func main() {
	var num int
	fmt.Println("Please Insert a Number:")
	_, _ = fmt.Scanf("%d", &num)

	switch num {
	case 1:
		// Go语言中的case默认包括break，因此不需要写break！
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fallthrough //使用 fallthrough将会强制执行下一个case的代码
	case 3:
		fmt.Println(3)
		fallthrough //但一般不使用该关键字
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Mamba Out!")

	}
}
