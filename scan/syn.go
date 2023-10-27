package scan

type SynScanner struct {
	timeout int
	thread  int
}

// NewSynScanner 创建SynScanner
func NewSynScanner(timeout int, thread int) *SynScanner {
	return &SynScanner{
		timeout: timeout,
		thread:  thread,
	}
}

// Start 用来启用工作池
func (s *SynScanner) Start() error {

	return nil
}

// Scan 实现的syn扫描函数
func (s *SynScanner) Scan(ip string, port int, timeout int) {

}
