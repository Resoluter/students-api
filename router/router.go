package router

import (
	"github.com/Resoluter/students-api/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/students/{id}", middleware.GetStudent).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/students/", middleware.GetAllStudents).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/students/", middleware.CreateStudent).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/students/{id}", middleware.UpdateStudent).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/students/{id}", middleware.DeleteStudent).Methods("DELETE", "OPTIONS")

	return router
}
