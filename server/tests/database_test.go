package main

import (
	"github.com/Fhoust/Go-app/server/common"
	"github.com/Fhoust/Go-app/server/database"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDatabase(t *testing.T) {
	common.SetupENV()

	database.SetupDB()

	db := database.GetDB()

	var user User

	stmt, _ := db.Prepare("insert into users(name) values(?)")

	dbReturn, _ := stmt.Exec("New potato")

	id, _ := dbReturn.LastInsertId()

	db.QueryRow("select id, name from users where id = ?", id).Scan(&user.ID, &user.Name)

	assert.Equal(t, "New potato", user.Name)
	assert.Equal(t, id, int64(user.ID))
}
