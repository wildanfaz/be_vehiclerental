package interfaces

import (
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"github.com/wildanfaz/vehicle_rental/src/libs"
)

type VehiclesRepo interface {
	FindAllVehicles() (*models.Vehicles, error)
	SaveVehicle(body *models.Vehicle) (*models.Vehicle, error)
	ChangeVehicle(vars string, body *models.Vehicle) (*models.Vehicle, error)
	RemoveVehicle(vars string, body *models.Vehicle) (*models.Vehicle, error)
	FindVehicle(search string) (*models.Vehicles, error)
	VehicleDetail(id string)(*models.Vehicle, error)
	RatingVehicles(offset int) (*models.Vehicles, error)
	CategoryVehicles(typeVehicle string) (*models.Vehicles, error)
	Location(location string) (*models.Vehicles, error)
	Type(typeVehicle string) (*models.Vehicles, error)
	CheckId(vars string, body *models.Vehicle) error
}

type VehiclesService interface {
	GetAllVehicles() *libs.Resp
	AddVehicle(body *models.Vehicle) *libs.Resp
	UpdateVehicle(vars string, body *models.Vehicle) *libs.Resp
	DeleteVehicle(vars string, body *models.Vehicle) *libs.Resp
	SearchVehicle(search string) *libs.Resp
	PopularVehicles(offset int) *libs.Resp
	GetVehicleDetail(id string) *libs.Resp
	TypeVehicles(typeVehicle string) *libs.Resp
	SortLocation(location string) *libs.Resp
	SortType(typeVehicle string) *libs.Resp
}
