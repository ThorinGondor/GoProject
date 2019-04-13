package main

import "fmt"

//匿名函数：无函数名称

func main() {
	a := 10
	str := "Kimi"

	f1 := func() {
		fmt.Println("A", a)
		fmt.Println("str =", str)
	}

	type FuncType func() //该函数无参数也无返回值

	//声明变量
	var f2 FuncType

	f2 = f1
	f2()

	/**下面是真正使用匿名函数的一系列方法：**/
	//定义一个匿名函数，同时调用
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}() //ATTENTION: 该圆括号代表调用此匿名函数

	//定义带参数的匿名函数，同时调用
	func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}(1, 2)

	//定义有参有返回值的匿名函数，同时调用
	f4 := func(args ...int) (max, min int) {
		flagMax := args[0]
		flagMin := args[0]
		for _, data := range args {
			if data >= flagMax {
				flagMax = data
			}
			if data <= flagMin {
				flagMin = data
			}
		}
		max = flagMax
		min = flagMin
		return
	}
	max, min := f4(3, 8, 7, 6, 5, 9, 2, 3, 4, 5)
	fmt.Printf("MAX: %d, MIN: %d\n", max, min)
}
