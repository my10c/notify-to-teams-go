// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
//

package configurator

import (
	"fmt"
	"os"

	// local
    "vars"

	// on github	
	"github.com/my10c/packages-go/print"
	"github.com/BurntSushi/toml"
)

type (
	tomlConfig struct {
		Teams vars.TeamsConfig `toml:"teams"`
	}
)

var (
	Print = print.New()
)

func GetConfig() vars.TeamsConfig {
	var configValues tomlConfig
	var configured vars.TeamsConfig

	if _, err := toml.DecodeFile(vars.TeamsConfigFile, &configValues); err != nil {
		Print.PrintRed("Error reading the configuration file\n")
		fmt.Fprintln(os.Stderr, err)
        Print.PrintBlue("Aborting...\n")
		os.Exit(1)
	}
	if	len(configValues.Teams.WebHookUrl) == 0 ||
		len(configValues.Teams.MonitorUrl) == 0 {
		Print.PrintRed("Error reading the configuration file, some value are missing or is empty\n")
		Print.PrintBlue("Make sure webhook_url and monitor_url are set\n")
        Print.PrintBlue("Aborting...\n")
        os.Exit(1)
	}
	configured.WebHookUrl = configValues.Teams.WebHookUrl
	configured.MonitorUrl = configValues.Teams.MonitorUrl
	return configured	
}
