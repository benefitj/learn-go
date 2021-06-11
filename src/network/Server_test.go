package network

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("init...")
}

func TestStartUdpServer(t *testing.T) {
	fmt.Println("TestStartUdpServer...")
	StartUdpServer()
}

func TestStartTcpServer(t *testing.T) {
	fmt.Println("TestStartTcpServer...")
	StartTcpServer()
}
