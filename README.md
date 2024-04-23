# notify-to-teams
pipe message from nagios/naemon to a teams channel
It requires the message to be in certain format, check
how to setup your nagios/naemon notification with `notify-to-teams -s`

## usage

```
usage: notify-to-teams [-h|--help] [-v|--version] [-i|--info] [-t|--test]
                       [-s|--setup] [-S|--teams-config]

                       Simple script send a message to a teams webhook channel
                       via a piped message.

Arguments:

  -h  --help          Print help information
  -v  --version       Show version
  -i  --info          Show how to use notify-to-teams
  -t  --test          test mode, no message will be sent. Default: false
  -s  --setup         Show how to setup in nagios or naemon
  -S  --teams-config  Show how to setup the teams configuration file
```

# how to build

```
make
```
