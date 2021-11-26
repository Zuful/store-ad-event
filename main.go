package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	driverName string = "sqlite3"
	dbFilePath string = "./ad.db"
)

func main() {
	r := gin.Default()
	config := cors.Config{
		AllowOrigins: []string{"http://localhost:4200"},
	}
	r.Use(cors.New(config))

	r.POST("/event", handlerCreateEvent)
	r.GET("/event/:from/:to", handlerGetEvent)

	err := r.Run()
	CheckErr(err)
}
