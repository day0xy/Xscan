package utils

import (
	flag "github.com/spf13/pflag"
)

type FlagParams struct {
	Help     bool
	IpStr    string
	PortStr  string
	ScanType string

	//思路： 如果指定了值，就把结构体里的这两个值，赋值到connect和syn的结构体里去
	Thread  int
	TimeOut int

	//思路： bool变量来看是否启用ping,      nmap风格
	Pn bool
}

var Params = &FlagParams{}

func init() {
	//用bool来显示帮助信息
	flag.BoolVarP(&Params.Help, "help", "h", false, "show help message")
	flag.StringVar(&Params.ScanType, "type", "connect", "connect or syn scan type")
	flag.StringVarP(&Params.PortStr, "port", "p", "", "port to scan")
	flag.StringVarP(&Params.IpStr, "ip", "i", "", "ip to scan")
	flag.IntVarP(&Params.Thread, "thread", "t", 5000, "set thread value")
	flag.IntVar(&Params.TimeOut, "timeout", 3, "set timeout value")
	flag.BoolVar(&Params.Pn, "Pn", false, "disable ping")

}
