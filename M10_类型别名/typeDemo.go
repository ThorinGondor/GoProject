package main

import "fmt"

func main() {
	//给类型赋别名，例如int64类型，我们给其命名为bigint类型
	type bigint int64
	var a bigint
	fmt.Printf("%T\n", a)

	type (
		long int64
		char byte
	)

	var b long = 1109987492
	var ch char = 'h'

	fmt.Printf("%T %T\n", b, ch)

}
