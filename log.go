package main

import "fmt"

func getLogMessage(host string, serviceName string, method string, path string, error int, timeMillis int64, message string) (logMessage string) {
	logMessage = fmt.Sprintf("host=%s method=%s serviceName=%s "+
		"path=%s error=%d runTime=%d msg=%s",
		host, method, serviceName, path, error, timeMillis, message)
	return logMessage
}

func getConfigLogMessage(host string, serviceName string, method string, path string, analytics string, favs string, login string, error int, timeMillis int64, message string) (logMessage string) {
	logMessage = fmt.Sprintf("host=%s method=%s serviceName=%s "+
		"path=%s analytics=%s favs=%s login=%s error=%d runTime=%d msg=%s",
		host, method, serviceName, path, analytics, favs, login, error, timeMillis, message)
	return logMessage
}
