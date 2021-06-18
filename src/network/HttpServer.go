package network

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// 启动HTTP服务
func StartHttpServer() {

	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:8080", nil)

}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RemoteAddr: ", r.RemoteAddr)
	fmt.Println("Method: ", r.Method)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println("Header: ", r.Header)
	fmt.Println("Body: ", r.Body)

	w.Write([]byte("www.pingpong.com"))
}

// 测试WebSocket服务
func StartWebSocketServer() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
