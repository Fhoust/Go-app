package driver

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/zerepl/go-app/internal/domain/common"
	"log"
)

var (
	db *sql.DB
)

// SetupDB this function open a database connection
func SetupDB() *sql.DB {
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
	db.SetMaxOpenConns(8)
	log.Println("Successfully opened a new connection")

	err = db.Ping()
	if err != nil {
		log.Panic("Problems with database: ", err)
	}

	return db
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

	myDB.Exec("CREATE database IF NOT EXISTS goapp")
	myDB.Exec("USE goapp")
	myDB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id integer auto_increment,
		name varchar(80),
		PRIMARY KEY(id)
	)`)

	myDB.Close()
	log.Println("Finished migration")
}
