package main

import "fmt"

func TestBitOperation() {

	//testPointer()

	a := struct {
		x int
	}{}

	a.x = 101
	p := &a
	p.x += 200

	fmt.Println("p.x = ", p.x)

}

func testPointer() {
	var b1 = 01000011

	fmt.Println("^b1 = ", ^b1)
	fmt.Println("b1 &^ 0b0100010 = ", (b1 &^ 0100010))

	aa := 1
	//++aa
	//aa++

	//if aa++ > 1 {
	//}

	p := &aa
	*p++
	fmt.Println("p = ", aa)
	*p += 2
	fmt.Println("p = ", *p)

	//p++

	fmt.Println("aa == p", (aa == *p))

}
