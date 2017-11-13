package state

import (
	"sync"
	"time"
)

var (
	ports        = sync.Map{}
	lastCall     = sync.Map{}
	isUp         = sync.Map{}
	underStartup = sync.Map{}
)

func Port(author, projectName string) int {
	val, ok := ports.Load(author + "_" + projectName)
	if ok {
		return val.(int)
	}
	return 0
}

func SetPort(author, projectName string, port int) {
	ports.Store(author+"_"+projectName, port)
}

// Use this to remove the service after stopping the container
func MarkAsDown(author, projectName string) {
	isUp.Store(author+"_"+projectName, false)
}

// Use this to remove the service after stopping the container
func MarkAsUp(author, projectName string) {
	isUp.Store(author+"_"+projectName, true)
}

func IsUp(author, projectName string) bool {
	val, ok := isUp.Load(author + "_" + projectName)
	if !ok {
		return false
	}
	return val.(bool)
}

// Use this to remove the service after stopping the container
func MarkAsUnderStartup(author, projectName string) {
	underStartup.Store(author+"_"+projectName, true)
}

// Use this to remove the service after stopping the container
func MarkAsNotUnderStartup(author, projectName string) {
	underStartup.Store(author+"_"+projectName, false)
}

func IsUnderStartup(author, projectName string) bool {
	val, ok := underStartup.Load(author + "_" + projectName)
	if !ok {
		return false
	}
	return val.(bool)
}

// Returns wether you can use the service or has to start it
func LastCallIn(author, projectName string, d time.Duration) bool {
	lc, ok := lastCall.Load(author + "_" + projectName)
	if !ok {
		return false
	}
	return time.Now().Sub(lc.(time.Time)) < d
}

// Call this when proxying a request
func SetLastCall(author, projectName string) {
	lastCall.Store(author+"_"+projectName, time.Now())
}
