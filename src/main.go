package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//fmt.Printf("Hello World !")

	// 创建字符串类型的通道
	channel := make(chan string)
	// 创建producer，发送goroutine
	go producer("cat", channel)
	// 创建producer，发送goroutine
	go producer("dog", channel)
	// 消费
	customer(channel)
}

func producer(header string, channel chan <- string)  {
	var i = 0
	// 循环生成数据
	for {
		// 随机产生
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		// 等待1秒
		time.Sleep(time.Second)
		if i > 200 {
			break
		}

		i ++
	}
}

func customer(channel <- chan string)  {
	// 不停获取数据
	for {
		// 从通道取出数据
		message := <- channel
		// 打印
		fmt.Println("message: " + message + ", time: " + time.Now().UTC().String())
	}
}
