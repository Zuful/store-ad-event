package main

import (
	"strings"
	"testing"
)

func TestStringToInt64(t *testing.T) {
	var stringNumber string = "12345678910"

	err, _ := stringToInt64(stringNumber)

	if err != nil {
		t.Error("An error occured while trying to convert string to int64", err)
	}
}

func TestGetOsFromUserAgent(t *testing.T) {
	var userAgent string = strings.ToLower(`Mozilla/5.0 (X11; Linux x86_64) 
												AppleWebKit/537.36 (KHTML, like Gecko) 
												Chrome/95.0.4638.69 Safari/537.36`)
	var osList []string = []string{"linux", "mac", "windows", "android", "ios"}
	var os string = "other"
	var expectedOs string = "linux"

	for _, oneOs := range osList {
		if strings.Contains(userAgent, oneOs) {
			os = oneOs
			break
		}
	}

	if os != expectedOs {
		t.Error("The os value " + os + " isn't the one expected: " + expectedOs)
	}
}
