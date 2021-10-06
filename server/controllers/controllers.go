package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Fhoust/Go-app/server/database"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

// User struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Insert a new user in the database
func Insert(newUser string) int {
	//TODO check if the user already exists
	// TODO check received payload
	db := database.GetDB()

	stmt, _ := db.Prepare("insert into users(name) values(?)")

	dbReturn, _ := stmt.Exec(newUser)
	id, _ := dbReturn.LastInsertId()

	log.Printf("Inserted %s with the id %d\n", newUser, id)
	return int(id)
}

// UpdateUser info of one user
func UpdateUser(id int, name string) {
	//TODO check if the user exists
	db := database.GetDB()

	stmt, _ := db.Prepare("update users set name = ? where id = ?")
	stmt.Exec(name, id)

	log.Printf("Updated %d to %s\n", id, name)

}

// DeletePerId this function deletes one ID user from the database
func DeletePerId(id int) string {
	// TODO check if the user exists
	db := database.GetDB()

	var deadUser User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&deadUser.ID, &deadUser.Name)

	deleter, _ := db.Prepare("delete from users where id = ?")

	deleter.Exec(id)

	log.Printf("Deleted the id: %d\n", id)

	return deadUser.Name
}

// deleteAll this function deletes all ids of the database
func deleteAll(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, _ := db.Query("select * from users where id > ?", 0)
	deleter, _ := db.Prepare("delete from users where id = ?")

	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
		deleter.Exec(u.ID)
	}

	json, _ := json.Marshal(users)

	log.Println("All users was deleted...")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "All users was deleted\n%s",string(json))

}

// UserPerID return info of the one user
func UserPerID(id int) string {
	db := database.GetDB()

	var user User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&user.ID, &user.Name)

	log.Printf("Requested info about %d - %s", user.ID, user.Name)

	return user.Name
}

// AllUsers return all users in DB
func AllUsers() string {
	db := database.GetDB()

	rows, _ := db.Query("select * from users where id > ?", 0)
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}

	usersJson, _ := json.Marshal(users)

	log.Printf("Requested info about all users")

	return string(usersJson)
}
