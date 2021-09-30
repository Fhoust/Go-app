package main

import (
	"log"
	"net/http"
)

func main() {

	SetupENV()
	SetupDB()

	http.HandleFunc("/usuarios/", UsuarioHandler)
	http.HandleFunc("/delete/", DeleteUser)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/insert", Insert)
	log.Println("Initializing the app")
	log.Fatal(http.ListenAndServe(GetPort(), nil))
}
