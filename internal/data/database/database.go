package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"

	"github.com/zerepl/go-app/common"
)

var (
	db *sql.DB
)

// SetupDB this function open a database connection
func SetupDB() {
	migration()
	log.Println("Opening a new connection with database")
	dbURL, dbPassword, dbUser := common.GetDBVars()
	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPassword,
		Net:    "tcp",
		Addr:   dbURL + ":3306",
		DBName: "goapp",
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
	dbURL, dbPassword, dbUser := common.GetDBVars()

	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPassword,
		Net:    "tcp",
		Addr:   dbURL + ":3306",
	}

	myDB, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal("Not able to connected to the database")
		panic(err)
	}
	
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