package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marcelom47/go-gorm-restapi/db"
	"github.com/marcelom47/go-gorm-restapi/models"
	"github.com/marcelom47/go-gorm-restapi/routes"
)

func main() {
	db.DBConnection()

	db.DB.AutoMigrate(&models.User{}, &models.Task{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.GetUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
