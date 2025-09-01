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
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	//"context"
	// msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	// graphmodels "github.com/microsoftgraph/msgraph-sdk-go/models"

	// local
	"vars"
)

type (
	TeamJson struct {
		Text   string `json:"text"`
		MrkDwn bool   `json:"mrkdwn"`
	}
)

const (
	// teams emojie codes
	RedCircle    = "&#x1f534;"
	GreenCircle  = "&#x1f7e2;"
	YellowCircle = "&#x1f7e1;"
	ColorRed     = "darkred"
	ColorGreen   = "darkgreen"
	ColorYellow  = "yellow"
)

func SendMessage(msg string, config vars.TeamsConfig) bool {

	// Construct the message payload (e.g., an Adaptive Card JSON)
	messagePayload := map[string]interface{}{
		"type": "message",
		"attachments": []map[string]interface{}{
			{
				"contentType": "application/vnd.microsoft.card.adaptive",
				"contentUrl":  nil,
				"content": map[string]interface{}{
					"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
					"type":    "AdaptiveCard",
					"version": "1.2",
					"body": []map[string]interface{}{
						{
							"type":   "TextBlock",
							"MrkDwn": true,
							"text":   msg,
							"wrap":   false,
						},
					},
				},
			},
		},
	}
	jsonMsg, err := json.Marshal(messagePayload)
	if err != nil {
		os.Exit(3)
	}
	request, err := http.NewRequest("POST", config.WorkFlowUrl, bytes.NewBuffer(jsonMsg))
	if err != nil {
		fmt.Printf("Errored %v\n", err)
		os.Exit(3)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Errored %v\n", err)
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
	}()
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
	// 🟩 📡 🔴 🟢 🌕
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
	notification_msg := "no message provided"
	if len(data[3]) > 1 {
		notification_msg = strings.ReplaceAll(data[3], "ServiceState: ", "")
	}
	url := fmt.Sprintf("Monitorl URL: [%s](%s)", notification_host, config.MonitorUrl+notification_host)
	switch notification_type {
	case "Host:":
		// build the message for a host notification
		notification_state = strings.ReplaceAll(data[2], "HostState: ", "")
		if strings.Contains(data[1], "DOWN") {
			message = fmt.Sprintf("🔴 %s\n\n * host ***alert***\n * DOWN\n * %s\n",
				url, notification_state)
		}
		if strings.Contains(data[1], "UP") {
			message = fmt.Sprintf("🟢 %s\n\n * host ***recovered***\n * UP\n * %s\n",
				url, notification_state)
		}
	case "ServiceHost:":
		// <span style=color:darkgreen>
		notification_state = strings.ReplaceAll(data[1], "ServiceOutput: ", "")
		service_name = strings.ReplaceAll(data[2], "ServiceName: ", "")
		if strings.Contains(data[1], "OK") {
			message = fmt.Sprintf("🟢 %s\n\n * Service ***recovered***\n * %s\n * %s\n * %s\n",
				url, service_name, notification_state, notification_msg)
		} else {
			// <span style=color:darkred>
			message = fmt.Sprintf("🔴 %s\n\n * Service ***alert***\n * %s\n * %s\n * %s\n",
				url, service_name, notification_state, notification_msg)
		}
	default:
		re := regexp.MustCompile("status.*")
		url := fmt.Sprintf("[monitor home](%s)", re.ReplaceAllString(config.MonitorUrl, "main.cgi"))
		message = fmt.Sprintf("🌕 %s\n\n * unknown ***error*** occured\n * please check the monitor dashboard\n",
			url)
	}
	return message
}
