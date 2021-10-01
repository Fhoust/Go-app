package routes

import (
	"log"
	"net/http"

	"github.com/Fhoust/Go-app/common"
	"github.com/Fhoust/Go-app/controllers"
)

// Routes -> define endpoints
func Routes() {
	http.HandleFunc("/usuarios/", controllers.UsuarioHandler)

	log.Fatal(http.ListenAndServe(common.GetPort(), nil))
}