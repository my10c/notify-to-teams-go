// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
//

package help

import (
	"fmt"
	"os"

	// local
    "vars"

	// on github	
	"github.com/my10c/packages-go/print"
)

var (
	Print = print.New()
)

func Version() {
	Print.ClearScreen()
	Print.PrintCyan("Version: " + vars.MyVersion + "\n")
	Print.PrintPurple(vars.MyInfo + "\n")
	os.Exit(0)
}

func Info() {
	Print.ClearScreen()
	Print.PrintYellow(vars.MyProgname + " usage should not be use with any flags, unless you want to:\n")
	Print.PrintYellow("see this information 😈 (-i), test without actually sent the message 🤣 (-t),\n")
	Print.PrintYellow("- show to configure Naemon/Nagios (-s)\n")
	Print.PrintYellow("- show to configure teams configuratiom file (" + vars.TeamsConfigFile + ") (-S)\n")
	Print.PrintYellow("- see the version (-v)\n")
	Print.PrintPurple("It should be use with pipped data from a nagios or naemon command.\n")
	Print.PrintPurple("Example: /usr/bin/printf \"%s\" \"<some-data>\" | " +  vars.MyProgname + "\n\n")
	os.Exit(0)
}

func Setup() {
	Print.ClearScreen()
	Print.PrintGreen("**Do** note that the \\n are required! It is use to parse the message\n")
	Print.PrintBlue("The script depends of the variables passed and their order!\n")

	fmt.Printf("\n# notify-host-to-teams command definition\n")
	fmt.Printf("define command{\n")
	fmt.Printf("  command_name notify-host-to-teams\n")
	fmt.Printf("  command_line /usr/bin/printf \"%%b\"")
	fmt.Printf(" \"Host: $HOSTNAME$\\nHostOutput: $HOSTOUTPUT$\\nHostState: $HOSTSTATE$\\n\"")
	fmt.Printf(" | /usr/local/sbin/notify-to-teams 2>> /tmp/hosts_notification.log\n}\n\n")

	Print.PrintGreen("and\n")

	fmt.Printf("\n# notify-service-to-teams command definition\n")
	fmt.Printf("#define command{\n")
  	fmt.Printf("  command_name notify-service-by-teams\n")
  	fmt.Printf("  command_line  /usr/bin/printf \"%%b\"")
	fmt.Printf(" \"ServiceHost: $HOSTNAME$\\nServiceOutput: $SERVICEOUTPUT$\\n")
	fmt.Printf("ServiceName: $SERVICEDISPLAYNAME$\\nServiceState: $SERVICESTATE$\\n\"")
	fmt.Printf(" | /usr/local/sbin/notify-to-teams 2>> /tmp/services_notification.log\n}\n\n")

	os.Exit(0)
}

func SetupConfig() {
	Print.ClearScreen()
	Print.PrintYellow("The configuration file is: " + vars.TeamsConfigFile + "\n")
	Print.PrintPurple("\n[teams]\n")
	Print.PrintPurple("# these are required\n")
	Print.PrintPurple("webhook_url = \"company-webhook-url\"\n")
	Print.PrintPurple("monitor_url = \"company-monitor-url\"\n\n")
	Print.PrintBlue("webhook_url example: \"https://webhooks.your-domain.tld/WebhookHandler/<team-token>\"\n\n")
	Print.PrintBlue("monitor_url example: \"https://naemon.your-domain.tld/thruk/cgi-bin/status.cgi?host=\"\n\n")
	os.Exit(0)
}
