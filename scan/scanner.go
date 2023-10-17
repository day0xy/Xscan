package scan

/*
	Scanner接口
*/

// Scanner 接口
type Scanner interface {
	// Start 思路： 用来启动工作池，实现并发
	Start() error

	// Scan 思路: 根据scanType来调用不同scanner对象的Scan函数
	Scan() error

	// Print 思路： 打印信息
	Print()
}
