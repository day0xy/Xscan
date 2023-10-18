package utils

import (
	"github.com/day0xy/Xscan/vars"
	flag "github.com/spf13/pflag"
)

var Params = &vars.FlagParams{}

func init() {
	flag.StringVar(&Params.ScanType, "type", "connect", "connect or syn scan type")
	flag.StringVarP(&Params.PortStr, "port", "p", "", "port to scan")
	flag.StringVarP(&Params.IpStr, "ip", "i", "", "ip to scan")
	flag.IntVarP(&Params.Thread, "thread", "t", 5000, "set thread value")
	flag.IntVar(&Params.TimeOut, "timeout", 3, "set timeout value")
	//用bool来显示报错信息
	flag.BoolVarP(&Params.Help, "help", "h", false, "show help message")
	flag.BoolVar(&Params.Pn, "Pn", false, "disable ping")

}
