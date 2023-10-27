package main

import (
	"fmt"
	"github.com/day0xy/Xscan/utils"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	params := utils.Params
	if params.Help {
		flag.PrintDefaults()
	}

	ports, err := utils.ParsePort(params.PortStr)
	if err != nil {
		fmt.Println("error in main.go,          parse port error!")
	}

	ip, err := utils.ParseIP(params.IpStr)
	if err != nil {
		fmt.Println("error in main.go,          parse ip error!")
	}

}
