package users

import (
	"github.com/gorilla/mux"
	"github.com/wildanfaz/vehicle_rental/src/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	//** unused get all users
	// route.HandleFunc("", middleware.CheckAuth([]string{"Admin"}, ctrl.GetAllUsers)).Methods("GET")
	route.HandleFunc("", middleware.HandlerChain(middleware.CheckAuth("User", "Admin")).Then(ctrl.GetUser)).Methods("GET")
	route.HandleFunc("/{name}", middleware.HandlerChain(middleware.CheckAuth("User", "Admin")).Then(ctrl.GetUserByName)).Methods("GET")

	//** register
	route.HandleFunc("", ctrl.AddUser).Methods("POST")

	route.HandleFunc("/{name}", middleware.HandlerChain(middleware.CheckAuth("User")).Then(ctrl.UpdateUser)).Methods("PUT")
	route.HandleFunc("/{name}", middleware.HandlerChain(middleware.CheckAuth("Admin")).Then(ctrl.DeleteUser)).Methods("DELETE")
}
