package main

import (
	"fmt"
	"os"
	"time"
)

// TODO: Add some sort of log cleanup.

type Logger struct {
	service 	string
	path    	string
	servicePath string
}

// Service name should be class creating the log.
func NewLogger(service string) (logger *Logger) {
	logger = &Logger{
		service: service,
		path:    "temp/logs",
		servicePath: "temp/logs/" + service,
	}

	logger.initializeLogger()

	return
}

func (logger *Logger) initializeLogger() error {
	err := os.MkdirAll(logger.path, 0755)
	if err != nil {
		return err
	}

	// I'm thinking its a good idea to have well organized logs. So each service gets its own folder with logs separated by
	// info/debug/error + a "sequential" log that just shows everything in chronological order, then cumulative info/debug/error
	// logs + a cumulative sequential log.
	for _, path := range []string{"Info", "Debug", "Error", "Sequential"} {
		err = os.WriteFile(logger.path + "/logs.txt", []byte(fmt.Sprintf("[%s] | (initializeLogger): logger file has been created...", time.Now().String())), 0755)
	if err != nil {
		return err
	}
	}
	
	return nil
}

// The writeLog function will write the log everywhere it needs to be written in. So first the cumulative sequential,
// then the service sequential, then the cumulative categorized, then the service categorized.
func (logger *Logger) writeLog(message, category string) error {
	for _, path := range []string{} {

	}

	return nil
}

func (logger *Logger) InfoLog(message string) {

}

func (logger *Logger) DebugLog(message string) {

}

func (logger *Logger) ErrorLog(message string) {

}
