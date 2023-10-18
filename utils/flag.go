package utils

import (
	"github.com/day0xy/Xscan/vars"
	flag "github.com/spf13/pflag"
)

var params = &vars.Params{}

func init() {
	flag.StringVar(&params.ScanType, "type", "connect", "connect or syn scan type")
	flag.StringVarP(&params.PortStr, "port", "p", "", "port to scan")
	flag.StringVarP(&params.IpStr, "ip", "i", "", "ip to scan")
	flag.IntVarP(&params.Thread, "thread", "t", 5000, "set thread value")
	flag.IntVar(&params.TimeOut, "timeout", 3, "set timeout value")
	//用bool来显示报错信息
	flag.BoolVarP(&params.Help, "help", "h", false, "show help message")
	flag.BoolVar(&params.Pn, "Pn", false, "disable ping")

}
