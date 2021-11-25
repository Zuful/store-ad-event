package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var (
	driverName string
	host       string
	port       string
	user       string
	password   string
	dbname     string

	psqlInfo string
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

func init() {
	err := godotenv.Load("db.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	driverName = "postgres"
	host = os.Getenv("CC_HOST")
	port = os.Getenv("CC_PORT")
	user = os.Getenv("CC_USER")
	password = os.Getenv("CC_PASSWORD")
	dbname = os.Getenv("CC_DB_NAME")

	psqlInfo = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		host, port, user, password, dbname)
}
