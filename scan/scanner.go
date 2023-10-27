package scan

import (
	"context"
	"fmt"
)

// Scanner 接口
type Scanner interface {
	// Start 思路： 用来启动工作池，实现并发
	Start(ctx context.Context, ip []string, port []int) (<-chan Result, <-chan error)

	// Scan 思路: 根据scanType来创建不同的scanner,调用不同scanner对象的Scan函数
	Scan(ctx context.Context, jobChan <-chan PortJob, resultChan chan<- Result, errChan chan<- error)
}

// CreateScanner 根据类型创建scanner
func CreateScanner(scanType string, timeout int, thread int) (Scanner, error) {
	switch scanType {
	case "connect":
		return NewConnectScanner(timeout, thread), nil
	case "syn":
		fmt.Println("under construction!")

	}
	return nil, fmt.Errorf("unknown scan type %s", scanType)
}
