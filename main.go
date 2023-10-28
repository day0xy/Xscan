package main

import (
	"context"
	"fmt"
	"github.com/day0xy/Xscan/scan"
	"github.com/day0xy/Xscan/utils"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	params := utils.Params
	if params.Help {
		flag.PrintDefaults()
	} else {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		ports, err := utils.ParsePort(params.PortStr)
		if err != nil {
			fmt.Println("error in main.go,          parse port error!")
		}

		ips, err := utils.ParseIP(params.IpStr)
		if err != nil {
			fmt.Println("error in main.go,          parse ip error!")
		}

		scanner, err := scan.CreateScanner(params.ScanType, params.TimeOut, params.Thread)
		results, errs := scanner.Start(ctx, ips, ports)
		scan.PrintResults(results, errs)
	}

}
