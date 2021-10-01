package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Fhoust/Go-app/database"
)

// Usuario struct
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega a funcao adequada
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New access in /usuarios/")
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorID(w, r, id)
	case r.Method == "GET":
		usuarioTodos(w, r)
	case r.Method == "DELETE" && id > 0:
		deletePerId(w, r, id)
	case r.Method == "DELETE":
		deleteAll(w, r)
	case r.Method == "UPDATE" && id > 0:
		update(w, r, id)
	case r.Method == "POST":
		insert(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Wrong method\n")
	}
}

func insert(w http.ResponseWriter, r *http.Request) {
	//TODO check if the user already exists
	db := database.GetDB()

	stmt, _ := db.Prepare("insert into usuarios(nome) values(?)")

	var newUser Usuario

	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(rawBody), &newUser)

	dbReturn, _ := stmt.Exec(newUser.Nome)
	id, _ := dbReturn.LastInsertId()
	w.Write([]byte("Inserted\n"))
	log.Printf("Inserted %s with the id %d\n", newUser.Nome, id)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s was added with the id: %d", newUser.Nome, id)
}


func update(w http.ResponseWriter, r *http.Request, id int) {
	//TODO check if the user exists
	db := database.GetDB()

	var newUser Usuario

	rawBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	json.Unmarshal([]byte(rawBody), &newUser)

	w.Write([]byte("Updated\n"))

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec(newUser.Nome, id)

	log.Printf("Updated %d to %s\n", newUser.ID, newUser.Nome)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%d was updated to the name %s", newUser.ID, newUser.Nome)
}

// deletePerId this function deletes one ID user from the database
func deletePerId(w http.ResponseWriter, r *http.Request, id int) {
	// TODO check if the user exists
	db := database.GetDB()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)

	deleter, _ := db.Prepare("delete from usuarios where id = ?")

	deleter.Exec(id)

	log.Printf("Deleted the id: %d\n", id)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s with the id %d was deleted", u.Nome, u.ID)
}

// deleteAll this function deletes all ids of the database
func deleteAll(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, _ := db.Query("select * from usuarios where id > ?", 0)
	deleter, _ := db.Prepare("delete from usuarios where id = ?")

	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var u Usuario
		rows.Scan(&u.ID, &u.Nome)
		usuarios = append(usuarios, u)
		deleter.Exec(u.ID)
	}

	json, _ := json.Marshal(usuarios)

	log.Println("All users was deleted...")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "All users was deleted\n%s",string(json))

}

func usuarioPorID(w http.ResponseWriter, r *http.Request, id int) {
	db := database.GetDB()

	var u Usuario
	db.QueryRow("select id, nome from usuarios where id = ?", id).Scan(&u.ID, &u.Nome)

	json, _ := json.Marshal(u)

	log.Printf("Requested info about %d - %s", u.ID, u.Nome)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	rows, _ := db.Query("select * from usuarios where id > ?", 0)
	defer rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nome)
		usuarios = append(usuarios, usuario)
	}

	json, _ := json.Marshal(usuarios)

	log.Printf("Requested info about all users")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
