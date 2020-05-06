package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"time"
)

var serviceName = "login"
var APP_NAME = "perfeng_search"
var hostName = "https://loginingo.herokuapp.com/"
var loaderIOPath = "loaderio-f55c136af99151b21f7ddb26dd774a61.txt"
var loaderIOUrlPath = "/loaderio-f55c136af99151b21f7ddb26dd774a61.txt"

type User struct {
	userid     int64
	username   string
	email      string
	age        int64
	isloggedin string
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
	router.POST("/", func(c *gin.Context) {
		fmt.Printf("URL to store: \n")
	})
	router.POST("/login", login)
	router.POST("/logout", logout)
	router.StaticFile(loaderIOUrlPath, loaderIOPath)

	router.Run(":" + port)
}

func login(c *gin.Context) {
	start := time.Now()
	id := 1
	var event *analyticsEvent

	db, err := getDbConn()
	if err != nil {
		logMessage := getLogMessage(hostName, serviceName, "GET", "/login", 1, time.Since(start).Nanoseconds()/1000000, "\"db connection error\"")
		log.Println(logMessage)
		log.Println(err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": "internal server error",
		})
		//event = getEvent("/search", time.Since(start).Nanoseconds()/1000, "500", false, start)
		//go postEvent(event)
		return
	}

	defer db.Close()
	var users []User
	sqlStatement := `SELECT * FROM users WHERE age = $1;`
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		logMessage := getLogMessage(hostName, serviceName, "GET", "/login", 1, time.Since(start).Nanoseconds()/1000000, "\"error querying db\"")
		log.Println(logMessage)
		log.Println(err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": "internal server error",
		})
		event = getEvent("/search", time.Since(start).Nanoseconds()/1000, "500", false, start)
		go postEvent(event)
		return
	}

	for rows.Next() {
		var user User
		queryErr := rows.Scan(&user.age, &user.email, &user.userid, &user.username, &user.isloggedin)
		switch queryErr {
		case sql.ErrNoRows:
			c.JSON(200, gin.H{
				"success": "true",
				"results": "[]",
			})
			event = getEvent("/search", time.Since(start).Nanoseconds()/1000, "200", true, start)
			go postEvent(event)
			return
		case nil:
			users = append(users, user)
			println(user.username)
			c.JSON(200, gin.H{
				"success": "true",
				"message": "login succeed ",
			})
		default:
			logMessage := getLogMessage(hostName, serviceName, "GET", "/search", 1, time.Since(start).Nanoseconds()/1000000, "\"error querying db\"")
			log.Println(logMessage)
			log.Println(err)
			c.JSON(500, gin.H{
				"success": "false",
				"message": "internal server error",
			})
			event = getEvent("/search", time.Since(start).Nanoseconds()/1000, "500", false, start)
			go postEvent(event)
			return
		}
	}
}

func logout(c *gin.Context) {
	start := time.Now()
	id := 1
	var event *analyticsEvent

	db, err := getDbConn()
	if err != nil {
		logMessage := getLogMessage(hostName, serviceName, "GET", "/login", 1, time.Since(start).Nanoseconds()/1000000, "\"db connection error\"")
		log.Println(logMessage)
		log.Println(err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": "internal server error",
		})
		//event = getEvent("/search", time.Since(start).Nanoseconds()/1000, "500", false, start)
		//go postEvent(event)
		return
	}

	defer db.Close()
	var users []User
	sqlStatement := `SELECT * FROM users WHERE age = $1;`
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		logMessage := getLogMessage(hostName, serviceName, "GET", "/login", 1, time.Since(start).Nanoseconds()/1000000, "\"error querying db\"")
		log.Println(logMessage)
		log.Println(err)
		c.JSON(500, gin.H{
			"success": "false",
			"message": "internal server error",
		})

		event = getEvent("/search", time.Since(start).Nanoseconds()/1000, "500", false, start)
		go postEvent(event)
		return
	}

	for rows.Next() {
		var user User
		queryErr := rows.Scan(&user.age, &user.email, &user.userid, &user.username, &user.isloggedin)
		switch queryErr {
		case sql.ErrNoRows:
			c.JSON(200, gin.H{
				"success": "true",
				"results": "[]",
			})
			event = getEvent("/search", time.Since(start).Nanoseconds()/1000, "200", true, start)
			go postEvent(event)
			return
		case nil:
			users = append(users, user)
			println(user.username)
			c.JSON(200, gin.H{
				"success": "true",
				"message": "logout succeed ",
			})
		default:
			logMessage := getLogMessage(hostName, serviceName, "GET", "/search", 1, time.Since(start).Nanoseconds()/1000000, "\"error querying db\"")
			log.Println(logMessage)
			log.Println(err)
			c.JSON(500, gin.H{
				"success": "false",
				"message": "internal server error",
			})
			event = getEvent("/search", time.Since(start).Nanoseconds()/1000, "500", false, start)
			go postEvent(event)
			return
		}
	}
}
