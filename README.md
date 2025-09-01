# notify-to-teams

## piped message
pipe message from nagios/naemon to a slack channel
It requires the message to be in certain format, check
how to setup your nagios/naemon notification with `notify-to-slack -s`


## direct message
use the `-m` flags and direct post a message to slack


## configuration
use the `-S` to see how to setup the configuration file

## usage

```
usage: notify-to-teams_Darwin_arm64 [-h|--help] [-c|--configFile "<value>"]
                                    [-v|--version] [-i|--info] [-t|--test]
                                    [-s|--setup] [-S|--teams-config]
                                    [-m|--message "<value>" [-m|--message
                                    "<value>" ...]] [-q|--quiet]

                                    Simple script send a message to a teams
                                    channel via a piped message, by the givev
                                    message on the command line.

Arguments:

  -h  --help          Print help information
  -c  --configFile    Configuration file to be use. Default:
                      /usr/local/etc/teams/teams.conf
  -v  --version       Show version
  -i  --info          Show how to use notify-to-teams_Darwin_arm64
  -t  --test          test mode, no message will be sent. Default: false
  -s  --setup         Show how to setup in nagios or naemon
  -S  --teams-config  Show how to setup the teams configuration file
  -m  --message       Message to be sent between double quotes or single
                      quotes, implies no stdin reading
  -q  --quiet         Quiet mode. Default: false

```

# how to build

```
make
```
