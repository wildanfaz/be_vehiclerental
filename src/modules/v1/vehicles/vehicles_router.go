package vehicles

import (
	"github.com/gorilla/mux"
	"github.com/wildanfaz/vehicle_rental/src/middleware"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/api/v1/vehicles").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", ctrl.GetAllVehicles).Methods("GET")
	route.HandleFunc("/detail/{id}", ctrl.GetVehicleDetail).Methods("GET")
	route.HandleFunc("/search", ctrl.SearchVehicle).Methods("GET")
	route.HandleFunc("/popular/{offset}", ctrl.PopularVehicles).Methods("GET")

	//**add vehicle with upload image
	// route.HandleFunc("", middleware.HandlerChain(middleware.CheckAuth("Admin"), middleware.UploadFileImage).Then(ctrl.AddVehicle)).Methods("POST")
	route.HandleFunc("", middleware.HandlerChain(middleware.CheckAuth("Admin"), middleware.CloudinaryAddImg).Then(ctrl.AddVehicle)).Methods("POST")

	route.HandleFunc("/{vehicle_id}", middleware.HandlerChain(middleware.CheckAuth("Admin")).Then(ctrl.UpdateVehicle)).Methods("PUT")
	route.HandleFunc("/{vehicle_id}", middleware.HandlerChain(middleware.CheckAuth("Admin")).Then(ctrl.DeleteVehicle)).Methods("DELETE")

	//**test chain midlleware
	route.HandleFunc("/v", middleware.HandlerChain(middleware.CheckAuth("User"), middleware.Hello, middleware.UploadFileImage).Then(ctrl.AddVehicle)).Methods("POST")

}
