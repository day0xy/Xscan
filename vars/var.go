package vars

type FlagParams struct {
	IpStr    string
	PortStr  string
	ScanType string
	Help     bool
	//思路： 如果指定了值，就把结构体里的这两个值，赋值到connect和syn的结构体里去
	Thread  int
	TimeOut int
	//思路： bool变量来看是否启用ping
	Pn bool
}
