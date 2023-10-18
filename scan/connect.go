package scan

import (
	"fmt"
	"net"
	"time"
)

type ConnectScanner struct {
	timeout time.Duration
	thread  int
}

// NewConnectScanner 初始化ConnectScanner
func NewConnectScanner(timeout time.Duration, thread int) *ConnectScanner {
	return &ConnectScanner{
		timeout: timeout,
		thread:  thread,
	}
}

// Start 用来启用工作池
func (s *ConnectScanner) Start() error {

	return nil
}

// Scan connect scan函数
func (s *ConnectScanner) Scan(ip string, port int, timeout time.Duration) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), timeout)
	if err != nil {
		fmt.Println("connect net.Dial error!")
	}
	conn.Close()

}
