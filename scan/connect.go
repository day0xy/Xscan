package scan

type ConnectScanner struct {
	timeout int
	thread  int
}

// NewConnectScanner 初始化ConnectScanner
func NewConnectScanner(timeout int, thread int) *ConnectScanner {
	return &ConnectScanner{
		timeout: timeout,
		thread:  thread,
	}
}

func (s *ConnectScanner) Start() error {

	return nil
}

func (s *ConnectScanner) Scan() error {

	return nil

}
