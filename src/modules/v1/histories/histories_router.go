package histories

import (
	"github.com/gorilla/mux"
	"github.com/wildanfaz/vehicle_rental/src/modules/v1/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/histories").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", middleware.HandlerChain(middleware.CheckAuth("User", "Admin")).Then(ctrl.GetAllHistories)).Methods("GET")
	route.HandleFunc("/search", middleware.HandlerChain(middleware.CheckAuth("User", "Admin")).Then(ctrl.SearchHistory)).Methods("GET")
	route.HandleFunc("", middleware.HandlerChain(middleware.CheckAuth("User", "Admin")).Then(ctrl.AddHistory)).Methods("POST")
	route.HandleFunc("/{history_id}", middleware.HandlerChain(middleware.CheckAuth("Admin")).Then(ctrl.UpdateHistory)).Methods("PUT")
	route.HandleFunc("/{history_id}", middleware.HandlerChain(middleware.CheckAuth("Admin")).Then(ctrl.DeleteHistory)).Methods("DELETE")
}
