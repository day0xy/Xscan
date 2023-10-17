package utils

type Params struct {
	ipStr    string
	portStr  string
	scanType string

	//思路： 如果指定了值，就把结构体里的这两个值，赋值到connect和syn的结构体里去
	thread  int
	timeOut int
	//思路： bool变量来看是否启用ping
	Pn bool
}

func init() {

}
