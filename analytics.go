package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type analyticsEvent struct {
	Method    string `json:"method"`
	Path      string `json:"path"`
	TimeMilis int64  `json:"processingTimeInMiliseconds"`
	Response  string `json:"responseCode"`
	Service   string `json:"serviceName"`
	Success   bool   `json:"success"`
	Timestamp string `json:"timestamp"`
	Username  string `json:"username"`
}

var analyticsHost = fmt.Sprint(os.Getenv("ANALYTICS_URL"), "/saveEdr")
var storeAnalytics = os.Getenv("STORE_ANALYTICS") == "true"

func getEvent(path string, timeMillis int64, response string, success bool, timestamp time.Time) *analyticsEvent {
	e := analyticsEvent{
		Method:    "GET",
		Path:      path,
		TimeMilis: timeMillis,
		Response:  response,
		Service:   "search",
		Success:   success,
		Timestamp: strconv.Itoa(int(timestamp.UTC().Unix())) + "000",
		Username:  "",
	}
	return &e
}

func postEvent(e *analyticsEvent) {
	storeAnalytics = os.Getenv("STORE_ANALYTICS") == "true"
	if !storeAnalytics {
		return
	}
	jsonEvent, err := json.Marshal(e)
	if err != nil {
		log.Println("internal error:", err)
		return
	}

	req, err := http.NewRequest("POST", analyticsHost, bytes.NewBuffer(jsonEvent))
	if err != nil {
		log.Println("POST to /saveEdr: internal error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("POST to /saveEdr: updated analytics with status:", resp.Status)
}
