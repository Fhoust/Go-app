package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"

	"github.com/Fhoust/Go-app/server/common"
)

var (
	db *sql.DB
	firstRun = true
)

// SetupDB this function open a database connection
func SetupDB() {

	log.Println("Opening a new connection with database")
	dbURL, dbPassword, dbUser := common.GetDBVars()
	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPassword,
		Net:    "tcp",
		Addr:   dbURL + ":3306",
		DBName: "goapp",
	}
	myDB, _ := sql.Open("mysql", cfg.FormatDSN())

	if firstRun {
		migration()
	}

	if err := myDB.Ping(); err != nil {
		log.Panicf("Not able to connected to the database: %v", err)
	} else {
		log.Println("Successfully opened a new connection")
	}

	db = myDB

	firstRun = false
}

// GetDB returns the DB instance
func GetDB() *sql.DB {
	err := db.Ping()
	if err != nil {
		SetupDB()
		log.Panicf("Problems with database: %v", err)
	}
	return db
}

// CloseDB close the connection between app and database
func CloseDB() {
	db.Close()
}

// migration prepare the database for the app
func migration() {
	dbURL, dbPassword, dbUser := common.GetDBVars()

	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPassword,
		Net:    "tcp",
		Addr:   dbURL + ":3306",
	}

	myDB, _ := sql.Open("mysql", cfg.FormatDSN())

	if err := myDB.Ping(); err != nil {
		log.Fatal("Not able to connected to the database: ", err)
	} else {
		log.Println("Successfully opened a new connection")
	}

	log.Println("Starting migration")
	
	myDB.Exec("create database if not exists goapp")
	myDB.Exec("use goapp")
	myDB.Exec(`create table if not exists users (
		id integer auto_increment,
		name varchar(80),
		PRIMARY KEY(id)
	)`)

	myDB.Close()
	log.Println("Finished migration")
}