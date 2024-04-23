// BSD 3-Clause License
//
// Copyright (c) 2023, © Badassops LLC / Luc Suryo
// All rights reserved.
//
// Version	:	0.1
//

package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"io"
	"strings"
	"time"
	"regexp"
	"net/http"

	// local
	"vars"
)

type (
	TeamJson struct {
		Text  string `json:"text"`
		MrkDwn bool `json:"mrkdwn"`
	}
)

const (
	RedCircle    = "&#x1f534;"
	GreenCircle  = "&#x1f7e2;"
	YellowCircle = "&#x1f7e1;"
	ColorRed     = "darkred"
	ColorGreen   = "darkgreen"
	ColorYellow  = "yellow"
)

func SendMessage(msg string, config vars.TeamsConfig) bool {
 	// create a new connection
	jsonData := TeamJson{
		Text: msg,
		MrkDwn: true,
	}
	jsonMsg, err := json.Marshal(jsonData)
	if err != nil {
		os.Exit(3)
	}
	request, err := http.NewRequest("POST", config.WebHookUrl, bytes.NewBuffer(jsonMsg))
	if err != nil {
		os.Exit(3)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	return true
}

func GetMessage(config vars.TeamsConfig) string {
	var msg string
	result := make(chan string, 1)
	go func() {
		result <- getMessage(config)
	} ()
	select {
		// we should get data within 2 seconds
		// otherwise we exit
		case <-time.After(2 * time.Second):
			os.Exit(3)
		case msg = <-result:
		break
	}
	return msg
}

func getMessage(config vars.TeamsConfig) string {
	//
	// 🟩 📡 🔴 🟢 
	//
	//	# HOSTSTATE				UP DOWN UNREACHABLE
	//	# SERVICESTATE			OK WARNING UNKNOWN CRITICAL <-- no longer needed, SERVICEOUTPUT have the info needed
	//	# SERVICEOUTPUT			[OK WARNING UNKNOWN CRITICAL] some text
	//	# SERVICEDISPLAYNAME	Alias of the check
	//
	var message string
	var notification_type string
	var notification_host string
	var notification_state string
	var service_name string

	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	// we know we need at least 60 chars, exit if we get less then 50 chars
	if len(string(stdin)) < 50 {
		os.Exit(3)
	}
	data := strings.Split(string(stdin), "\n")
	notification_type = strings.Split(data[0], " ")[0]
	notification_host = strings.Split(data[0], " ")[1]
	url := fmt.Sprintf("[%s](%s)", notification_host, config.MonitorUrl, notification_host)
	switch notification_type {
		case "Host:":
			// build the message for a host notification
			notification_state = strings.ReplaceAll(data[2], "HostState: ", "")
			if strings.Contains(data[1],"DOWN") {
				message = fmt.Sprintf("%s %s\n\n * host <span style=color:%s>***alert***</span>\n * DOWN\n * %s\n",
				 RedCircle, url, ColorRed, notification_state)
			}
			if strings.Contains(data[1], "UP") {
				message = fmt.Sprintf("%s %s\n\n * host <span style=color:%s>***recovered***</span>\n * UP\n * %s\n",
				 GreenCircle, url, ColorGreen, notification_state)
			}
		case "ServiceHost:":
			// <span style=color:darkgreen>
			notification_state = strings.ReplaceAll(data[1], "ServiceOutput: ", "")
			service_name = strings.ReplaceAll(data[2], "ServiceName: ", "")
			if strings.Contains(data[1],"OK") {
				message = fmt.Sprintf("%s %s\n\n * Service <span style=color:%s>****recovered***</span>\n * %s\n * %s\n",
				 GreenCircle, url, ColorGreen, service_name, notification_state)
			} else {
				// <span style=color:darkred>
				message = fmt.Sprintf("%s %s\n\n * Service <span style=color:%s>***alert***</span>\n * %s\n * %s\n",
				 RedCircle, url, ColorRed, service_name, notification_state)
			}
		default:
			re := regexp.MustCompile("status.*")
			url := fmt.Sprintf("[monitor home](%s)", re.ReplaceAllString(config.MonitorUrl, "main.cgi"))
			message = fmt.Sprintf("%s %s\n\n * unknown <span style=color:%s>***error***</span> occured\n * please check the monitor dashboard\n",
				YellowCircle, url, ColorYellow)
	}
	return message
}
