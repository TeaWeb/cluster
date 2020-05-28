package configs

import (
	"regexp"
	"strings"
	"sync"
)

var fileLocker = &sync.Mutex{}
var idRegexp = regexp.MustCompile(`^\w+$`)

func lock() {
	fileLocker.Lock()
}

func unlock() {
	fileLocker.Unlock()
}

func mapVars(s string, variables map[string]string) string {
	for k, v := range variables {
		s = strings.Replace(s, "${"+k+"}", v, -1)
	}
	return s
}
