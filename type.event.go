package main

import (
	"database/sql"
)

type Event struct {
	Id        int    `json:"id"`
	Type      string `json:"type"`
	UserAgent string `json:"user_agent"`
	Ip        string `json:"ip"`
	Timestamp int64  `json:"timestamp"`
	OsName    string `json:"os_name"`
}

func createEvent(event Event) error {
	database, err := sql.Open(driverName, dbFilePath)
	CheckErr(err)
	defer database.Close()

	statement, err := database.Prepare(`INSERT INTO ad_event(	type , user_agent, ip, timestamp , os_name)
											  VALUES (?, ?, ?, ?, ?)`)
	_, err = statement.Exec(event.Type, event.UserAgent, event.Ip, event.Timestamp, event.OsName)
	CheckErr(err)

	return err
}

func getEventList(from, to int64) (error, []Event) {
	database, err := sql.Open("sqlite3", dbFilePath)
	CheckErr(err)
	defer database.Close()

	var eventList []Event = make([]Event, 0)
	var sqlQuery string = `SELECT id, type, user_agent, ip, timestamp, os_name
						   FROM ad_event WHERE timestamp BETWEEN ? AND ? ORDER BY timestamp ASC LIMIT 200`

	result, err := database.Query(sqlQuery, from, to)

	for result.Next() {
		var event Event
		err = result.Scan(&event.Id, &event.Type, &event.UserAgent, &event.Ip, &event.Timestamp, &event.OsName)

		eventList = append(eventList, event)
	}

	return err, eventList
}

func getAllGroupsEventList(eventList []Event, allGroups []string) map[string]map[string][]Event {
	var groupedEventList map[string]map[string][]Event = make(map[string]map[string][]Event)

	for _, oneGroup := range allGroups {
		groupedEventList[oneGroup] = groupEventList(eventList, oneGroup)
	}

	return groupedEventList
}

func groupEventList(eventList []Event, groupBy string) map[string][]Event {
	var groupedEventList map[string][]Event = make(map[string][]Event)

	for _, oneEvent := range eventList {
		if groupBy == "type" {
			groupedEventList[oneEvent.Type] = append(groupedEventList[oneEvent.Type], oneEvent)
		} else if groupBy == "os" {
			groupedEventList[oneEvent.OsName] = append(groupedEventList[oneEvent.OsName], oneEvent)
		} else {
			groupedEventList["all"] = append(groupedEventList["all"], oneEvent)
		}
	}

	return groupedEventList
}
