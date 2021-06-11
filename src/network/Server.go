package network

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func StartUdpServer() {
	host := flag.String("host", "127.0.0.1", "host")
	port := flag.String("port", "37", "port")

	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address:", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
	}

	defer conn.Close()

	for {
		handleClient(conn)
	}

}

func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}

	daytime := time.Now().Unix()
	fmt.Println(n, remoteAddr)
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(daytime))
	conn.WriteToUDP(b, remoteAddr)
}

func StartTcpServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}
