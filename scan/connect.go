package scan

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"
)

type ConnectScanner struct {
	timeout int
	thread  int
	ctx     context.Context
}

type PortJob struct {
	ip   string
	port int
}

type PortState uint8

// 枚举类型
const (
	PortOpen PortState = iota
	PortClosed
)

type Result struct {
	Host  string
	Ports map[int]PortState
}

// NewConnectScanner 初始化ConnectScanner
func NewConnectScanner(timeout int, thread int) *ConnectScanner {
	return &ConnectScanner{
		timeout: timeout,
		thread:  thread,
	}
}

func (s *ConnectScanner) Start(ctx context.Context, ip []string, port []int) (<-chan Result, <-chan error) {
	jobChan := make(chan PortJob)
	resultChan := make(chan Result)
	errChan := make(chan error)

	var wg sync.WaitGroup

	// 创建并启动协程
	for i := 0; i < s.thread; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.Scan(ctx, jobChan, resultChan, errChan)
		}()
	}

	// 分发端口扫描任务
	go func() {
		defer close(jobChan)
		for _, ip := range ip {
			for _, port := range port {
				select {
				case <-ctx.Done():
					return
				case jobChan <- PortJob{ip: ip, port: port}:
				}
			}
		}
	}()

	// 等待所有的协程完成
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	return resultChan, errChan
}

func (s *ConnectScanner) Scan(ctx context.Context, jobChan <-chan PortJob, resultChan chan<- Result, errChan chan<- error) {
	results := make(map[string]map[int]PortState)

	for job := range jobChan {
		state, err := s.ScanPort(job.ip, job.port)
		if err != nil {
			errChan <- err
			continue
		}

		if _, ok := results[job.ip]; !ok {
			results[job.ip] = make(map[int]PortState)
		}
		results[job.ip][job.port] = state
	}

	for ip, ports := range results {
		result := Result{Host: ip, Ports: ports}

		select {
		case <-ctx.Done():
			return
		case resultChan <- result:
		}
	}
}

func (s *ConnectScanner) ScanPort(ip string, port int) (PortState, error) {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), time.Duration(s.timeout)*time.Second)
	if err != nil {
		return PortClosed, nil
	}
	conn.Close()
	return PortOpen, nil
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
				fmt.Printf("%s:%d is %s\n", result.Host, port, status)
			}
		case err, ok := <-errs:
			if !ok {
				return
			}
			fmt.Println("Error:", err)
		}
	}
}
