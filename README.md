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
usage: notify-to-teams [-h|--help] [-c|--configFile "<value>"] [-v|--version]
                       [-i|--info] [-t|--test] [-s|--setup] [-S|--teams-config]
                       [-m|--message "<value>" [-m|--message "<value>" ...]]
                       [-q|--quiet]

                       Simple script send a message to a teams channel via a
                       piped message or by the given the message on the command
                       line.

Arguments:

  -h  --help          Print help information
  -c  --configFile    Configuration file to be use. Default:
                      /usr/local/etc/teams/teams.conf
  -v  --version       Show version
  -i  --info          Show how to use notify-to-teams
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

## NOTES
- required a teams workflow url
- [Retirement of Office 365 connectors within Microsoft Teams](https://devblogs.microsoft.com/microsoft365dev/retirement-of-office-365-connectors-within-microsoft-teams/)
- search for [create a teams workflow to post to a channel](https://www.google.com/search?q=create+a+teams+workflow+to+post+to+a+channel&sca_esv=d7ac497bd0642435&rlz=1C5CHFA_enUS1166US1166&ei=pNe1aMr4FsOwqtsPzeOYiAY&oq=create+a+teams+workflow+to+post+&gs_lp=Egxnd3Mtd2l6LXNlcnAiIGNyZWF0ZSBhIHRlYW1zIHdvcmtmbG93IHRvIHBvc3QgKgIIADIFECEYoAEyBRAhGKABMgUQIRigATIFECEYoAEyBRAhGKABMgUQIRirAjIFECEYqwJIgzBQhwdYpSBwAHgBkAEAmAFroAHhB6oBAzMuN7gBA8gBAPgBAZgCCqACygfCAgQQABhHwgIGEAAYFhgewgILEAAYgAQYhgMYigXCAggQABiABBiiBMICCBAAGKIEGIkFwgIFECEYnwWYAwDiAwUSATEgQIgGAZAGCJIHAzIuOKAHjT6yBwMxLji4B8QHwgcFMC40LjbIBx4&sclient=gws-wiz-serp)
