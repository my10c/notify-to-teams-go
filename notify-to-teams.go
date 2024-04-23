// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
//

package main

import (
	"fmt"
	"os"

	// local
	conf "configurator"
	args "getargs"
	msg "message"
)

var (
    test bool
)

func main() {
	test = args.GetArgs()

	config := conf.GetConfig()
	teamsMsg := msg.GetMessage(config)
	if test {
		fmt.Printf("%s\n%v\n", teamsMsg)
	} else {
		msg.SendMessage(teamsMsg, config)
	}
	os.Exit(0)
}
