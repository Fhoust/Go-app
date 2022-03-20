package controllers

import (
	"log"
	"net/http"

	"github.com/zerepl/go-app/common"
	"github.com/zerepl/go-app/controllers"
)

// Routes -> define endpoints
func Routes() {
	http.HandleFunc("/users/", controllers.userHandler)

	log.Println("Initializing the app")

	log.Fatal(http.ListenAndServe(common.GetPort(), nil))
}

// UserHandler analyze the request and delegate the proper function
func userHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("New access in /users/")
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		usersPerID(w, r, id)
	case r.Method == "GET":
		allUsers(w, r)
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