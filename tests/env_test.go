package main

import (
	"github.com/zerepl/go-app/common"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("DB_URL", "mydb.com")
	os.Setenv("DB_USER", "dbUser")
	os.Setenv("DB_PASS", "dbPass")
	os.Setenv("PORT", "8080")

	defer os.Unsetenv("DB_URL")
	defer os.Unsetenv("DB_USER")
	defer os.Unsetenv("DB_PASS")
	defer os.Unsetenv("PORT")

	common.SetupENV()

	dbURL, dbPassword, dbUser := common.GetDBVars()

	assert.Equal(t, "mydb.com", dbURL)
	assert.Equal(t, "dbUser", dbUser)
	assert.Equal(t, "dbPass", dbPassword)

	port := common.GetPort()


	assert.Equal(t, ":8080", port)

}
