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
