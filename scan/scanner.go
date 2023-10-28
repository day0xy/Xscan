package scan

import (
	"context"
	"fmt"
)

// Scanner 接口
type Scanner interface {
	Start(ctx context.Context, ip []string, port []int) (<-chan Result, <-chan error)
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
