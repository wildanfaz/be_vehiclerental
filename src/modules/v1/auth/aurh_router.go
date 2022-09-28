package auth

import (
	"github.com/gorilla/mux"
	"github.com/wildanfaz/vehicle_rental/src/modules/v1/users"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/auth").Subrouter()

	repo := users.NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", ctrl.SignIn).Methods("POST")
}
