package main

import (
	"fmt"
)

func TestStruct() {

	// 测试结构体
	testStruct()

	// 切片
	testSlice()

}

// 切片
func testSlice() {

}

func testStruct() {
	var book = Book{
		name:  "Go语言程序设计",
		price: 30.12,
	}
	fmt.Println("书名: ", book.name, "，价格: ", book.price)

	book = Book{"Java编程语言", 40.3}
	fmt.Println("书名: ", book.name, "，价格: ", book.price)

	fmt.Println(book)

	var goBook = Book{}
	goBook.name = "Go语言实战"
	goBook.price = 23.3
	fmt.Println("地址: ", &goBook.name, ", goBook: ", goBook)

	var goBook2 = Book{}
	goBook2.name = "Go语言实战"
	goBook2.price = 23.3
	fmt.Println("地址: ", &goBook2.name, ", goBook: ", goBook2)

	fmt.Println("go1 == go2  ==>: ", goBook.name == goBook2.name)

	var goBook3 = funcStruct(goBook2)
	fmt.Println("地址: ", &goBook3.name, ", goBook: ", goBook3)

	fmt.Println("go2 == go3  ==>: ", goBook2 == goBook3)

}

func funcStruct(book Book) Book {
	book.name = "Java实战"
	book.price = 345.0
	return book
}

// book
type Book struct {
	// 名称
	name string
	// 加个
	price float32
}
