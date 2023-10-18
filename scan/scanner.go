package scan

import (
	"fmt"
	"time"
)

// Scanner 接口
type Scanner interface {
	// Start 思路： 用来启动工作池，实现并发
	Start() error

	// Scan 思路: 根据scanType来创建不同的scanner,调用不同scanner对象的Scan函数
	Scan(ip string, port int, timeout time.Duration)

	// Print 思路： 打印信息
}

// CreateScanner 根据类型创建scanner
func CreateScanner(scanType string, timeout time.Duration, thread int) (Scanner, error) {
	switch scanType {
	case "connect":
		return NewConnectScanner(timeout, thread), nil
	case "syn":
		return NewSynScanner(timeout, thread), nil
	}
	return nil, fmt.Errorf("unknown scan type %s", scanType)
}
