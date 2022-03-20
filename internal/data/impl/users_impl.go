package repositories

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/zerepl/go-app/database"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) spi.UserRepository {
	return &userRepository{db: db}
}


func insert(w http.ResponseWriter, r *http.Request) {
	//TODO check if the user already exists
	// TODO check received payload
	db := database.GetDB()

	stmt, _ := db.Prepare("insert into users(name) values(?)")

	var u model.User

	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(rawBody), &u)

	dbReturn, _ := stmt.Exec(u.Name)
	id, _ := dbReturn.LastInsertId()
	w.Write([]byte("Inserted\n"))
	log.Printf("Inserted %s with the id %d\n", u.Name, id)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s was added with the id: %d", u.Name, id)
}


func update(w http.ResponseWriter, r *http.Request, id int) {
	//TODO check if the user exists
	db := database.GetDB()

	var oldUser model.User

	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(rawBody), &oldUser)

	w.Write([]byte("Updated\n"))

	stmt, _ := db.Prepare("update users set name = ? where id = ?")
	stmt.Exec(oldUser.Name, id)

	log.Printf("Updated %d to %s\n", oldUser.ID, oldUser.Name)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%d was updated to the name %s", id, oldUser.Name)
}

// deletePerId this function deletes one ID user from the database
func deletePerId(w http.ResponseWriter, r *http.Request, id int) {
	// TODO check if the user exists
	db := database.GetDB()

	var deadUser model.User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&deadUser.ID, &deadUser.Name)

	deleter, _ := db.Prepare("delete from users where id = ?")

	deleter.Exec(id)

	log.Printf("Deleted the id: %d\n", id)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s with the id %d was deleted", deadUser.Name, deadUser.ID)
}

// deleteAll this function deletes all ids of the database
func deleteAll(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, _ := db.Query("select * from users where id > ?", 0)
	deleter, _ := db.Prepare("delete from users where id = ?")

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
		deleter.Exec(u.ID)
	}

	json, _ := json.Marshal(users)

	log.Println("All users was deleted...")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "All users was deleted\n%s",string(json))

}

func usersPerID(w http.ResponseWriter, r *http.Request, id int) {
	db := database.GetDB()

	var user model.User
	db.QueryRow("select id, name from users where id = ?", id).Scan(&user.ID, &user.Name)

	json, _ := json.Marshal(user)

	log.Printf("Requested info about %d - %s", user.ID, user.Name)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, _ := db.Query("select * from users where id > ?", 0)
	defer rows.Close()

	var users []model.User

	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name)
		users = append(users, u)
	}

	json, _ := json.Marshal(users)

	log.Printf("Requested info about all users")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
