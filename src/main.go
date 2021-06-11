package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {

	//TestMap()
	//TestArray()
	//TestSlice()

	//TestStruct2()
	//TestMethod()
	//TestJson()

	//TestSwitch()
	//TestSelect()
	//TestFor()
	//TestRange()
	//TestFunc()
	//TestClosure()
	//TestDefer()

	TestUserMethod()
}

// 测试Map
func TestMap() {
	mapTemp := map[string]string{
		"小李": "湖南",
		"小刘": "湖北",
	}
	mapTemp["小王"] = "北京"
	mapTemp["小崔"] = "天津"

	fmt.Println("length: ", len(mapTemp))
	value, ok := mapTemp["小李"]
	fmt.Println("value: ", value, ", ok: ", ok)

	if v, ok := mapTemp["小李"]; ok {
		fmt.Println("v ==>: ", v)
	}

	for k, v := range mapTemp {
		fmt.Println("k = ", k, ", v = ", v)
	}

}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// 测试数组
func TestArray() {

	arr1 := [2][3]int{{1, 2, 3}, {9, 4, 5}}
	fmt.Println(arr1)

	fmt.Println("&arr1 ==>: ", &arr1)
	fmt.Printf("arr1: %p\n", &arr1)

	fmt.Println("----------------- printArr-----------------------")
	var arr2 [5]int
	printArr(&arr2)
	fmt.Println("--------------")
	arr3 := [...]int{1, 3, 5, 20, 2212}
	printArr(&arr3)
	fmt.Println("--------------")
	fmt.Println("arr3: ", arr3)
	fmt.Println("----------------- printArr-----------------------")

}

// 测试切片
func TestSlice() {

	var i1 []int
	if i1 == nil {
		fmt.Println("空")
	} else {
		fmt.Println("满")
	}

	i2 := []int{}
	fmt.Println("i2.len = ", len(i2))

	i3 := make([]int, 10)
	fmt.Println(i1, i2, i3)

	// 4.初始化赋值
	var i4 []int = make([]int, 0, 0)
	fmt.Println("i4 =>: ", i4)
	i4 = append(i4, 1)
	fmt.Println("i4   cap=", cap(i4), " len=", len(i4))

	i5 := [5]int{2, 5, 9, 20, 1}
	fmt.Println(i5)

	fmt.Println("1:3 ==>: ", i5[1:3])
	fmt.Println(":3 ==>: ", i5[:3])
	fmt.Println("2: ==>: ", i5[2:])
	fmt.Println(": ==>: ", i5[:])
	fmt.Println("i5   cap=", cap(i5), " len=", len(i5))

	i6 := [...]int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	fmt.Println("i6   cap=", cap(i6), " len=", len(i6))

}

func TestStruct2() {

	// 别名
	type Byte int8
	type MyByte = int8

	const a Byte = 8
	const b MyByte = 8

	fmt.Println("a = ", a, ", b = ", b)
	fmt.Printf("a = %T, b = %T \n", a, b)

	var p Person
	p.name = "礼拜天"
	p.age = 7
	p.sex = "难"
	fmt.Println("person: ", p)
	fmt.Printf("person: %T, &p = %p", p, &p)

	var user struct {
		Name string
		Age  int
	}
	user.Name = "星期五"
	user.Age = 10
	fmt.Println("user =>: ", user)
	fmt.Printf("user =>: %T\n", user)

	fmt.Println("\n-----------")

	var p2 = new(Person)
	p2.age = 20
	p2.sex = "牛"
	p2.name = "周一"
	fmt.Printf("p2. ==>: %T\n", p2)
	fmt.Printf("p2.. ==>: %#\n", p2)

	fmt.Println("\n-----------")

	p3 := &Person{}
	p3.name = "黑色星期四"
	p3.sex = "难"
	p3.age = 4
	fmt.Printf("p3. ==>: %T\n", p3)
	fmt.Printf("p3.. ==>: %#\n", p3)

	fmt.Println("\n-----------")

	p4 := Person{"开心", "女", 20}
	fmt.Printf("p4. ==>: %T\n", p4)
	fmt.Printf("p4.. ==>: %#\n", p4)
	fmt.Println()
	fmt.Printf("p4.name ==>: %p\n", &p4.name)
	fmt.Printf("p4.age ==>: %p\n", &p4.age)
	fmt.Printf("p4.sex ==>: %p\n", &p4.sex)

	p5 := newPerson("喜欢数学的女孩", 20, "女")
	fmt.Println("p5 ==>: ", p5)

}

type Person struct {
	name string
	sex  string
	age  int
}

func newPerson(name string, age int, sex string) *Person {
	return &Person{
		name: name,
		age:  age,
		sex:  sex,
	}
}

func SayHello(p *Person) {
	fmt.Printf("%s say hello...  %p \n", p.name, &p)
}

// 测试方法
func TestMethod() {

	//func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
	//   函数体
	//}

	p := newPerson("小丁同学", 30, "男")
	p.Chaserainbows()

	fmt.Println("&p ==>: ", &p)
	SayHello(p)

	s := &Student{
		DeptNode: "1222",
		Person: &Person{
			name: "夏冬梅",
			age:  20,
			sex:  "女",
		},
	}
	fmt.Println("s ==>: ", s)
	fmt.Println("&s.Person ==>: ", &s.Person)
	SayHello(s.Person)
}

func (p Person) Chaserainbows() {
	fmt.Printf("%s做了个白日梦！\n", p.name)
}

type Student struct {
	DeptNode string // 学号
	*Person
}

func TestJson() {
	s := &Student{
		DeptNode: "1222",
		Person: &Person{
			name: "夏冬梅",
			age:  20,
			sex:  "女",
		},
	}
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)

	data, err = json.Marshal(s.Person)
	fmt.Printf("json:%s\n", data)
	fmt.Printf("name:%s\n", s.name)
}

func TestSwitch() {

	grade := "B"

	var r string

	switch grade {
	case "A":
		r = "优"
	case "B":
		r = "良"
	case "C":
		r = "中"
	case "D":
		r = "差"
	default:
		r = "糟糕"
	}
	fmt.Println("r =>: ", r)

	var x interface{}
	switch x.(type) {
	case string:
		fmt.Println("string...")
	case nil:
		fmt.Println("nil...")
	case float32:
	case float64:
		fmt.Println("float...")
	case int8:
	case int16:
	case int32:
	case int64:
		fmt.Println("int...")
	default:
		fmt.Println("未知...")
	}

	n := rand.Int31n(10902)
	switch {
	case n >= 0 && n < 10:
		fmt.Println("1.  n = ", n)
	case n >= 10 && n < 100:
		fmt.Println("2.  n = ", n)
	case n >= 100 && n < 1000:
		fmt.Println("3.  n = ", n)
	default:
		fmt.Println("4.  n = ", n)
	}

}

// 测试select语句
func TestSelect() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is cloed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

	ch := make(chan int, 5)
	data := 0
	select {
	case ch <- data:
		fmt.Println("data: ", data)
	default:
		// discard
	}

	data = 10

}

func TestFor() {

	numbers := [6]int{1, 2, 3, 5, 9, 2}
	for i, x := range numbers {
		fmt.Printf("第 %d 位的 x 的值 = %d\n", i, x)
	}

	//for {
	//	// 无线循环
	//}

}

func TestRange() {
	//Golang range类似迭代器操作，返回 (索引, 值) 或 (键, 值)。

	str := "abcdefg"
	for i := range str {
		fmt.Printf("index[%d] = %s\n", i, str[i])
	}
	for _, v := range str {
		println(v)
	}

	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		println("key: ", k, " value: ", v)
	}
}

// 测试函数
func TestFunc() {
	s1 := test2(func() int {
		return 1 + 2
	})
	fmt.Printf("s1: %d\n", s1)

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 22)
	fmt.Printf("s2: %s\n", s2)

	x, y := 12, 44
	sx, sy := Swap(&x, &y)
	fmt.Printf("sx: %d, sy: %d\n", sx, sy)

	TestVariableParameter(10, 20, 3, 643)

	TestInterface(func() {})

	fmt.Println("add(10, 22) = ", add(10, 22))
	fmt.Println("addWithDefer(10, 22) = ", addWithDefer(10, 22))

}

func test2(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func Swap(x, y *int) (int, int) {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
	return *x, *y
}

func TestVariableParameter(args ...int) {
	fmt.Printf("len: %d, %d \n", len(args), args)
}

func TestInterface(args ...interface{}) {
	for i, v := range args {
		fmt.Println("i: ", i, " v: ", v)
	}
}

func add(x, y int) (z int) {
	{
		var z = x + y
		return z
	}
}

func addWithDefer(x, y int) (z int) {
	defer func() {
		z += 102
	}()
	z = x + y
	return z
}

// 测试闭包
func TestClosure() {

	a := ClosureA()
	a()
	a()
	a()

	b := ClosureA()
	b()

	a()

	fmt.Println("ClosureAdd------------")

	ca := ClosureAdd(10)
	fmt.Println("ca =>: ", ca(8))
	fmt.Println("ca =>: ", ca(7))
	fmt.Println("ca =>: ", ca(9))

	fmt.Println("ClosureCal------------")

	add, sub := ClosureCal(5)
	fmt.Println("add =>: ", add(8), ", sub =>: ", sub(7))
	fmt.Println("add =>: ", add(5), ", sub =>: ", sub(2))

	fmt.Printf("Factorial of %d is %d\n", 5, TestFactorial(5))

}

func ClosureA() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println("i = ", i)
		return i
	}
	return b
}

func ClosureAdd(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func ClosureCal(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

// 测试阶乘（递归）
func TestFactorial(i int) int {
	if i > 1 {
		return i * TestFactorial(i-1)
	}
	return 1
}

// 测试Defer
func TestDefer() {
	//testBug1()
	//testBug2()
	//testBug3()
	//testBug4()
	TestDeferHttpGet()
}

func testBug1() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

type TestDeferType struct {
	name string
}

func (t *TestDeferType) Close() {
	fmt.Println(t.name, " closed")
}

func Close(t TestDeferType) {
	fmt.Println(t.name, " closed")
}

func testBug2() {
	ts := []TestDeferType{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		//defer t.Close()
		//defer Close(t)

		tt := t
		defer tt.Close()
	}
}

var lock sync.Mutex

func test() {
	lock.Lock()
	lock.Unlock()
}

func testdefer() {
	lock.Lock()
	defer lock.Unlock()
}

func testBug3() {
	func() {
		t1 := time.Now()
		for i := 0; i < 10000; i++ {
			test()
		}
		elapsed := time.Since(t1)
		fmt.Println("test elapsed: ", elapsed)
	}()
	func() {
		t1 := time.Now()

		for i := 0; i < 10000; i++ {
			testdefer()
		}
		elapsed := time.Since(t1)
		fmt.Println("testdefer elapsed: ", elapsed)
	}()
}

func testBug4() (i int) {
	i = 0
	defer func() {
		fmt.Println("i ==>: ", i)
	}()
	return 2
}

func TestDeferHttpGet() {
	res, err := http.Get("http://192.168.1.198:8081/swagger-ui.html")
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		panic(err)
	}

	var data []byte
	size, _ := res.Body.Read(data)
	fmt.Println("size ==>: ", size)
	fmt.Println(string(data))
}

type User struct {
	Name string
	Age  int
}

func (self *User) Say(word string, change bool) {
	fmt.Printf("%s say: %s\n", self.Name, word)
	if change {
		self.Name = "王五"
	}
}

func (self User) Say2(word string) {
	fmt.Printf("%s say: %s\n", self.Name, word)
	self.Name = "王五"
}

type Manager struct {
	User
}

func (self *Manager) Say(word string) {
	fmt.Printf("%s say: %s\n", self.Name, word)
}

type Integer struct {
	int
}

func TestUserMethod() {

	// 值传递
	u := &User{"张三", 20}
	u.Say("世界，你好！", true)
	u.Say("世界，你好！", false)

	// 引用传递
	u2 := &User{"李四", 20}
	u2.Say2("世界，你好！")
	u2.Say2("世界，你好！")

	// 匿名方法和字段
	m := &Manager{}
	m.Name = "管理员"
	m.Age = 50
	m.User.Say("你好", false)
	m.Say("你好")

	i := Integer{10}
	fmt.Println("i = ", i)
	fmt.Println("i.int = ", i.int)

}
