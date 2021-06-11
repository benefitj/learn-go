package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func TestFileOperate() {

	start := time.Now()
	f, err := os.Open("D:/home/znsx/log/znsx-biz.log")

	if err != nil {
		fmt.Println("read fail")
		return
	}

	//// 通过块读取
	//content := readFileByChunkBuf(f)
	//// 通过IO工具
	//content := readFileByIoUtil(f)
	// 读取成reader
	content := readFileByReader(f)

	fmt.Println(content)
	//fmt.Println(len(content))

	// 写入文件
	writeFile(content, strings.Split(f.Name(), ".")[0]+".txt")

	fmt.Println("speed: ", time.Now().Sub(start).Milliseconds())

}

/// 从读取
func readFileByIoUtil(f *os.File) string {
	chunk, err := ioutil.ReadFile(f.Name())
	if err != nil {
		fmt.Println("read filed", err)
		return "err"
	}
	return string(chunk)
}

/// 读取文件
func readFileByChunkBuf(f *os.File) string {
	defer f.Close()
	var chunk []byte
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("read fail", err)
			return ""
		}
		if n == 0 {
			break
		}
		chunk = append(chunk, buf[:n]...)
	}
	return string(chunk)
}

func readFileByReader(f *os.File) string {
	defer f.Close()
	reader := bufio.NewReader(f)
	var lines []string
	for {
		line, _, err := reader.ReadLine()
		if err != nil && err != io.EOF {
			panic(err)
		}

		if line == nil {
			break
		}

		lines = append(lines, string(line))
	}
	return strings.Join(lines, "\n")
}

// 写入文件
func writeFile(data string, filename string) {
	file := OpenFile(filename)
	defer file.Close()
	n, err := io.WriteString(file, data)
	fmt.Println("写入文件：", filename, len(data), n)

	//err := ioutil.WriteFile(file.Name(), []byte(data), 0666)
	//fmt.Println("写入文件：", filename, len(data))

	if err != nil {
		fmt.Println("write fail ", err)
		return
	}
}

// 打开文件
func OpenFile(filename string) *os.File {
	if !CheckFileExist(filename) {
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		return file
	} else {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		return file
	}
}

// 检查文件是否存在
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
