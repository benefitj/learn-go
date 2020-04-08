package main

import (
	"fmt"
)

/**
 * 基本语法
 */

var (
	vname1 = "ssssss"
	vname2 = "hhhhhh"
)

func main() {

	//var age int = 20
	fmt.Println("age: " + "age")

	// 布尔类型: var b bool = true // false
	// 数值类型: int, float32, float64
	// 字符串:

	// 派生类型：
	// 指针类型，数组，结构化，Channel类型，函数类型，切片类型，接口类型，Map类型

	var aUint8 uint8 = 9
	fmt.Println("aUint8:", aUint8)

	//var b byte = 0x01

	var aStr string = "Hello World"
	fmt.Println("str ==>:", aStr)

	// 没有初始化的变量会被赋予默认值
	var i int16
	var b bool
	fmt.Println(i, b)

	var v1, v2, v3 = 2, "ss", false
	fmt.Println("自动类型推断:", v1, v2, v3)

	fmt.Println("因式分解全局变量:", vname1, vname2)

	g, h := 145, "22"
	fmt.Println("因式分解局部变量:", g, h)

	// 值类型和引用类型
	// int/float/string/bool 都是值类型，
	// 引用变量存储的是值所在的地址或第一个字的位置
	// 这个地址称之为指针，

	// 交换两个变量的值
	exchangeA, exchangeB := 10, 222
	exchangeA, exchangeB = exchangeB, exchangeA
	fmt.Println(exchangeA, exchangeB)

	var str string
	str2, str := "sss22", "..."
	str3, str4 := "sss22", "sss22"
	println(str2, str)

	println("address: ", &str2, &str3, &str4)

	str3, str4 = str4, str3
	println("address: ", &str2, &str3, &str4)

	testConst()
	testConst2()

	i = 1000
	// 条件表达式
	if i/10 == 100 {
		println("呵呵")
	}

	// 循环
	testFor()

	// 传递引用
	testReference()

	// 数组
	testArray()

}

func testArray() {
	println("\n---------------------------\n")
	var array = [][]int{{1, 2, 3, 5}, {4, 5, 3}}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			print(array[i][j], " ")
		}
		println()
	}
	println("\n---------------------------\n")
}

const CONST_A = "常量"

const (
	FEMALE  = 0
	MALE    = 1
	UNKNOWN = 2
)

func testConst() {
	const LENGTH = 10
	const WIDTH = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为: %d\n", area)

	fmt.Println("a: ", a, ", b: ", b, ", c: ", c)

	//FEMALE, MALE, UNKNOWN;

}

func testConst2() {

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)

}

// 测试循环
func testFor() {
	for j := 0; j < 100; j++ {
		if j/10 == 0 {
			print("j ==>: ", j)
		}
	}
	println()

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	println(sum)

	strings := []string{"google", "runoob"}
	for index, str := range strings {
		println(index, " ==>: ", str)
	}

	numbers := []int{1, 2, 3, 4, 5, 6}
	for index, num := range numbers {
		println("第", index, "位的值位: ", num)
	}

}

func swap(x *int, y *int) {
	//var temp int
	//temp = *x
	//*x = *y
	//*y = temp

	*x, *y = *y, *x
}

// 测试传递引用
func testReference() {
	x, y := 100, 1220
	swap(&x, &y)
	print("x = ", x, ", y = ", y)
}
