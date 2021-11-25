package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func handlerCreateEvent(c *gin.Context) {
	var code int = http.StatusOK
	var event Event

	err := c.Bind(&event)
	if err != nil {
		code = http.StatusInternalServerError
	}

	event.Ip = c.GetHeader("X-FORWARDED-FOR")
	event.UserAgent = c.GetHeader("User-Agent")
	event.OsName = getOsFromUserAgent(event.UserAgent)

	err = createEvent(event)
	if err != nil {
		code = http.StatusInternalServerError
	}

	c.JSON(code, err)
}

func handlerGetEvent(c *gin.Context) {
	var code int = http.StatusOK
	err, from := stringToInt64(c.Param("from"))
	if err != nil {
		code = http.StatusInternalServerError
	}

	err, to := stringToInt64(c.Param("to"))
	if err != nil {
		code = http.StatusInternalServerError
	}

	var allGroups []string = make([]string, 0)
	var allGroupsEventList map[string]map[string][]Event

	err, eventList := getEventList(from, to)
	if err != nil {
		code = http.StatusInternalServerError
	}

	if c.Query("groupBy") != "" {
		allGroups = strings.Split(c.Query("groupBy"), ",")
	} else {
		allGroups = append(allGroups, "*")
	}

	allGroupsEventList = getAllGroupsEventList(eventList, allGroups)

	c.JSON(code, allGroupsEventList)
}
