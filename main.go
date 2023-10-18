package main

import (
	"github.com/day0xy/Xscan/utils"
	flag "github.com/spf13/pflag"
)

func main() {
	flag.Parse()
	params := utils.Params
	if params.Help {
		flag.PrintDefaults()
	}

}
