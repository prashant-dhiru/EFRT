package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/spf13/viper"
)

type JiraTaskResponseType struct {
	Total  int `json:"total"`
	Issues []struct {
		Key    string `json:"key"`
		Fields struct {
			Summary string `json:"summary"`
		} `json:"fields"`
	} `json:"issues"`
}

type JsonPayloadToLogEffort struct {
	TimeSpentSeconds int    `json:"timeSpentSeconds"`
	Comment          string `json:"comment"`
}

var fetchAllActiveTaskPayload string = `{
	"jql": "assignee = currentUser() AND status = 'In Progress'",
	"startAt": 0,
	"maxResults": 20,
	"fields": [
		"id",
		"key",
		"summary"
	]
}`

func GetAllActiveTask() JiraTaskResponseType {

	jsonBody := []byte(fetchAllActiveTaskPayload)
	bodyReader := bytes.NewReader(jsonBody)

	jira_server := viper.Get("JIRA_SERVER")
	search_api := viper.Get("JIRA_API.SEARCH")
	url := fmt.Sprintf("%s%s", jira_server, search_api)

	jira_token := viper.Get("JIRA_ACCESS_TOKEN")
	bearer := fmt.Sprintf("Bearer %s", jira_token)

	contentType := "application/json"

	req, _ := http.NewRequest("POST", url, bodyReader)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", contentType)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	} else if resp.StatusCode == 401 {
		log.Println("Unautherized request, please reconfigure your JIRA PAT.\n[ERROR] -", err)
	} else if resp.StatusCode != 200 {
		log.Println("Error on response, response code \n[ERROR] -", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var jiraResponse JiraTaskResponseType
	err = json.Unmarshal(body, &jiraResponse)
	if err != nil {
		log.Println("error is decoding JSON:", err)
	}
	return jiraResponse
}

func LogEfforts(issuesKey string, effort string, comment string) {
	effortInSec := 0

	if effort[len(effort)-1:] == "m" {
		effortInSec, _ = strconv.Atoi(effort[:len(effort)-1])
		effortInSec = effortInSec * 60
	} else if effort[len(effort)-1:] == "h" {
		effortInSec, _ = strconv.Atoi(effort[:len(effort)-1])
		effortInSec = effortInSec * 60 * 60
	} else if effort[len(effort)-1:] == "h" {
		effortInSec, _ = strconv.Atoi(effort[:len(effort)-1])
		effortInSec = effortInSec * 60 * 60 * 8
	}

	jira_server := viper.Get("JIRA_SERVER")
	worklog_api := viper.Get("JIRA_API.WORKLOG")
	url := fmt.Sprintf("%s%s", jira_server, worklog_api)
	url = fmt.Sprintf(url, issuesKey)

	jira_token := viper.Get("JIRA_ACCESS_TOKEN")
	bearer := fmt.Sprintf("Bearer %s", jira_token)

	contentType := "application/json"

	jsonPayload, _ := json.Marshal(&JsonPayloadToLogEffort{
		TimeSpentSeconds: effortInSec,
		Comment:          comment,
	})
	bodyReader := bytes.NewReader(jsonPayload)
	// Create a new request using http
	req, _ := http.NewRequest("POST", url, bodyReader)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", contentType)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		log.Printf("%s effort logged for %s", effort, issuesKey)
	} else {
		log.Fatalf("error while logging effort for %s", issuesKey)
	}

}
