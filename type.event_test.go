package main

import (
	"fmt"
	"testing"
)

func TestCreateEvent(t *testing.T) {
	var event Event = Event{Type: "click", OsName: "linux", Ip: "127.0.0.1", Timestamp: 1637770227,
		UserAgent: `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) 
					Chrome/95.0.4638.69 Safari/537.36`}

	var err error = createEvent(event)

	if err != nil {
		t.Error("Error while trying to create an event", err)
	}
}

func TestGetEventList(t *testing.T) {
	var from int64 = 1637770000
	var to int64 = 1637771000

	err, _ := getEventList(from, to)

	if err != nil {
		t.Error("An error occured while trying to get the event list.", err)
	}
}

func TestGetAllGroupsEventList(t *testing.T) {
	var eventList []Event = []Event{
		{Type: "impression", OsName: "linux"},
		{Type: "impression", OsName: "mac"},
		{Type: "click", OsName: "linux"},
		{Type: "click", OsName: "mac"},
		{Type: "click", OsName: "android"},
	}
	var allGroups []string = []string{"*", "type", "os"}
	fmt.Println(getAllGroupsEventList(eventList, allGroups))
	for oneGroupName, oneGroupEventList := range getAllGroupsEventList(eventList, allGroups) {
		if oneGroupName == "*" && len(oneGroupEventList["all"]) != len(eventList) {
			t.Error("Size mismatch for the all group.")
		} else if oneGroupName == "type" {

			// checking the size of the impression type group
			if impressionEventList, ok := oneGroupEventList["impression"]; ok {
				if len(impressionEventList) != 2 {
					t.Error("Size mismatch for the impression group.")
				}
			}
			// checking the size of the click type group
			if clickEventList, ok := oneGroupEventList["click"]; ok {
				if len(clickEventList) != 3 {
					t.Error("Size mismatch for the click group.")
				}
			}

		} else if oneGroupName == "os" {
			// checking the size of the linux type group
			if linuxEventList, ok := oneGroupEventList["linux"]; ok {
				if len(linuxEventList) != 2 {
					t.Error("Size mismatch for the linux group.")
				}
			}
			// checking the size of the mac type group
			if macEventList, ok := oneGroupEventList["mac"]; ok {
				if len(macEventList) != 2 {
					t.Error("Size mismatch for the mac group.")
				}
			}
			// checking the size of the android type group
			if angroidEventList, ok := oneGroupEventList["android"]; ok {
				if len(angroidEventList) != 1 {
					t.Error("Size mismatch for the android group.")
				}
			}
		}
	}
}

func TestGroupEventList(t *testing.T) {
	var eventList []Event
	var groupBy string
	var groupedEventList map[string][]Event = make(map[string][]Event)

	for _, oneEvent := range eventList {
		if groupBy == "type" {
			groupedEventList[oneEvent.Type] = append(groupedEventList[oneEvent.Type], oneEvent)
		} else if groupBy == "os" {
			groupedEventList[oneEvent.OsName] = append(groupedEventList[oneEvent.OsName], oneEvent)
		}
	}
}
