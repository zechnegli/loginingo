package main

import (
	"database/sql"
	"fmt"
)

var (
	host     = "ec2-18-235-20-228.compute-1.amazonaws.com"
	port     = 5432
	user     = "szxecjisxxjtqi"
	password = "37d65ada41b263abafb16a28b160f8fa47f132f42272833168cae207d0cbb3f1"
	dbname   = "ddujk6ojbki5mb"
	sslmode  = "require"
)

func getDbConn() (db *sql.DB, err error) {
	// connect to database
	println(host)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err = sql.Open("postgres", psqlInfo)
	return
}
