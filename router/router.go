package router

import (
	"github.com/Resoluter/students-api/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/{id}", middleware.GetStudent).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/user/", middleware.GetAllStudents).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newuser", middleware.CreateStudent).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/{id}", middleware.UpdateStudent).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteuser/{id}", middleware.DeleteStudent).Methods("DELETE", "OPTIONS")

	return router
}
