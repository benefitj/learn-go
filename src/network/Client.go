package network

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func StartUdpClient() {
	var host = flag.String("host", "localhost", "host")
	var port = flag.String("port", "37", "port")

	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(""))
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}

	data := make([]byte, 4)
	_, err = conn.Read(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err)
		os.Exit(1)
	}

	t := binary.BigEndian.Uint32(data)
	fmt.Println(time.Unix(int64(t), 0).String())

	os.Exit(0)

}

func StartTcpClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
