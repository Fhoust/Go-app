package routes

import (
	"log"
	"net/http"

	"github.com/Fhoust/Go-app/controllers"
	common "github.com/Fhoust/Go-app/handlers"
)

// Routes -> define endpoints
func Routes() {
	http.HandleFunc("/usuarios/", controllers.UsuarioHandler())

	log.Fatal(http.ListenAndServe(common.GetPort(), nil))
}