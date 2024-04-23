// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
//

package vars

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"
)

type (
	TeamsConfig struct {
		WebHookUrl  string
		MonitorUrl  string
	}
)

var (
   MyVersion   = "0.0.1"
    now         = time.Now()
    MyProgname  = path.Base(os.Args[0])
    myAuthor    = "Luc Suryo"
    myCopyright = "Copyright 2023 - " + strconv.Itoa(now.Year()) + " ©Badassops LLC"
    myLicense   = "License 3-Clause BSD, https://opensource.org/licenses/BSD-3-Clause ♥"
    myEmail     = "<luc@badassops.com>"
    MyInfo      = fmt.Sprintf("%s (version %s)\n%s\n%s\nWritten by %s %s\n",
        MyProgname, MyVersion, myCopyright, myLicense, myAuthor, myEmail)
    MyDescription = "Simple script send a message to a teams webhook channel via a piped message."

	TeamsConfigFile = "/etc/teams.conf"
)
