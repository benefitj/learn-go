package network

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestStartHttpServer(t *testing.T) {
	// 启动HTTP服务
	StartHttpServer()
}

// 测试HTTP客户端
func TestStartHttpClient(t *testing.T) {
	resp, _ := http.Get("http://127.0.0.1:8080/go")
	defer resp.Body.Close()

	fmt.Println("Status: ", resp.Status)
	fmt.Println("StatusCode: ", resp.StatusCode)
	fmt.Println("Header: ", resp.Header)
	fmt.Println("ContentLength: ", resp.ContentLength)

	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println("res: ", res)
			break
		}
	}

}

// 测试WebSocket服务
func TestStartWebSocketServer(t *testing.T) {
	StartWebSocketServer()
}
