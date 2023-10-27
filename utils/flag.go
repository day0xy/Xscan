package utils

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

type FlagParams struct {
	Help     bool
	IpStr    string
	PortStr  string
	ScanType string
	Thread   int
	TimeOut  int
}

var Params = &FlagParams{}

func init() {
	//用bool来显示帮助信息
	fmt.Println("Example: Xscan -i 127.0.0.1 -p 3000-4000 --type connect[default:connect] ")
	flag.BoolVarP(&Params.Help, "help", "h", false, "show help message")
	flag.StringVar(&Params.ScanType, "type", "connect", "connect or syn scan type")
	flag.StringVarP(&Params.PortStr, "port", "p", "", "port to scan")
	flag.StringVarP(&Params.IpStr, "ip", "i", "", "ip to scan")
	flag.IntVarP(&Params.Thread, "thread", "t", 5000, "set thread value")
	flag.IntVar(&Params.TimeOut, "timeout", 3, "set timeout value")

}
