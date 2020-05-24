package main

import "fmt"

func main() {

	// 测试操作符
	testOperator()

}

// 操作符
func testOperator() {
	const v = 20

	var a byte = 20
	b := v + a
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", v, v)

	var c float32 = 1.2
	d := c + v
	fmt.Printf("%T, %v\n", d, d)

	// 位移操作符必须是无符号整数
	(func() {
		b := 23
		fmt.Printf("%T, %v\n", b, b)
		//x := 1 << b // b是有符号int，无法操作位移
		//println(x)
	})()

	aa := 1.0 << 3 // 包含常亮展开，1.0变为1，1.2无法常量展开，会报错

	fmt.Printf("%T, %v\n", aa, aa)

}
