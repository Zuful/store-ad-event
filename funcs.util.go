package main

import (
	"log"
	"strconv"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func stringToInt64(str string) (error, int64) {
	var output int64
	var err error

	if s, err := strconv.ParseInt(str, 10, 64); err == nil {
		output = s
	}

	return err, output
}

func getOsFromUserAgent(userAgent string) string {
	userAgent = strings.ToLower(userAgent)
	var osList []string = []string{"linux", "mac", "windows", "android", "ios"}
	var os string = "other"

	for _, oneOs := range osList {
		if strings.Contains(userAgent, oneOs) {
			os = oneOs
			break
		}
	}

	return os
}
