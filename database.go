package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

var (
	db *sql.DB
)

// SetupDB this function open a database connection
func SetupDB() {
	log.Println("Opening a new connection with database")
	dbURL, dbPassword, dbUser := GetDBVars()
	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPassword,
		Net:    "tcp",
		Addr:   dbURL + ":3306",
		DBName: "myapp",
	}
	myDB, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal("Not able to connected to the database")
		panic(err)
	}

	db = myDB

	db.SetMaxIdleConns(10)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(8)
	log.Println("Successfully opened a new connection")
	migration()
}

// GetDB returns the DB instance
func GetDB() *sql.DB {
	err := db.Ping()
	if err != nil {
		SetupDB()
		log.Panic("Problems with database: ", err)
	}
	return db
}

// CloseDB close the connection between app and database
func CloseDB() {
	db.Close()
}

// migration prepare the database for the app
func migration() {
	log.Println("Starting migration")
	db.Exec("create database if not exists myapp")
	db.Exec("use myapp")
	db.Exec(`create table if not exists usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY(id)
	)`)
}