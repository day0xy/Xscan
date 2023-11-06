package scan

import (
	"context"
	"fmt"
)

// Scanner 接口
type Scanner interface {
	Start(ctx context.Context, ip []string, port []int) (<-chan Result, <-chan error)
	Scan(ctx context.Context, jobChan <-chan PortJob, results map[string]map[int]PortState, errChan chan<- error)
	ScanPort(ctx context.Context, ip string, port int) (PortState, error)
}

// CreateScanner 根据类型创建scanner
func CreateScanner(scanType string, timeout int, thread int) (Scanner, error) {
	switch scanType {
	case "connect":
		return NewConnectScanner(timeout, thread), nil
	case "syn":
		return NewSynScanner(timeout, thread), nil
		//默认的话，就用os.Exit(1)来退出
	}
	return nil, fmt.Errorf("unknown scan type %s", scanType)
}

func PrintResults(results <-chan Result, errs <-chan error) {
	for {
		select {
		case result, ok := <-results:
			if !ok {
				return
			}

			fmt.Printf("Target %s:\n", result.Host)
			for port, state := range result.Ports {
				status := "closed"
				if state == PortOpen {
					status = "open"
				}
				fmt.Printf("%d is %s,\n", port, status)
			}
			fmt.Println()
		case err, ok := <-errs:
			if !ok {
				return
			}
			fmt.Println("Error:", err)
		}
	}
}
