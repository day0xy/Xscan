package scan

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
