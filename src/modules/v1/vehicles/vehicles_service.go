package vehicles

import (
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"github.com/wildanfaz/vehicle_rental/src/interfaces"
	"github.com/wildanfaz/vehicle_rental/src/libs"
)

type vehicles_service struct {
	repo interfaces.VehiclesRepo
}

func NewService(repo interfaces.VehiclesRepo) *vehicles_service {
	return &vehicles_service{repo}
}

func (svc *vehicles_service) GetAllVehicles() *libs.Resp {
	data, err := svc.repo.FindAllVehicles()

	if err != nil {
		return libs.Response(data, 400, "failed get data", err)
	}

	return libs.Response(data, 200, "success get data", nil)
}

func (svc *vehicles_service) AddVehicle(body *models.Vehicle) *libs.Resp {
	_, err := svc.repo.SaveVehicle(body)

	if err != nil {
		return libs.Response(nil, 400, "failed add data", err)
	}

	return libs.Response(nil, 201, "success add data", nil)
}

func (svc *vehicles_service) UpdateVehicle(vars string, body *models.Vehicle) *libs.Resp {
	if check := svc.repo.CheckId(vars, body); check != nil {
		return libs.Response(nil, 400, "failed update data", check)
	}

	data, err := svc.repo.ChangeVehicle(vars, body)

	if err != nil {
		return libs.Response(nil, 400, "failed update data", err)
	}

	return libs.Response(data, 200, "success update data", nil)
}

func (svc *vehicles_service) DeleteVehicle(vars string, body *models.Vehicle) *libs.Resp {
	if check := svc.repo.CheckId(vars, body); check != nil {
		return libs.Response(nil, 400, "failed delete data", check)
	}

	_, err := svc.repo.RemoveVehicle(vars, body)

	if err != nil {
		return libs.Response(nil, 400, "failed delete data", err)
	}

	return libs.Response(nil, 200, "success delete data", nil)
}

func (svc *vehicles_service) SearchVehicle(search string) *libs.Resp {
	data, err := svc.repo.FindVehicle(search)

	if err != nil {
		return libs.Response(data, 404, "failed search data", err)
	}

	return libs.Response(data, 200, "success search data", nil)
}

func (svc *vehicles_service) PopularVehicles(offset int) *libs.Resp {
	data, err := svc.repo.RatingVehicles(offset)

	if err != nil {
		return libs.Response(data, 400, "failed get data", err)
	}

	return libs.Response(data, 200, "success get data", nil)
}

func (svc *vehicles_service) GetVehicleDetail(id string) *libs.Resp {
	data, err := svc.repo.VehicleDetail(id)

	if err != nil {
		return libs.Response(data, 400, "failed get data", err)
	}

	return libs.Response(data, 200, "success get data", nil)
}

func (svc *vehicles_service) TypeVehicles(typeVehicle string) *libs.Resp {
	data, err := svc.repo.CategoryVehicles(typeVehicle)

	if err != nil {
		return libs.Response(data, 400, "failed get data", err)
	}

	return libs.Response(data, 200, "success get data", nil)
}

func (svc *vehicles_service) SortLocation(location string) *libs.Resp {
	data, err := svc.repo.Location(location)

	if err != nil {
		return libs.Response(data, 400, "failed get data", err)
	}

	return libs.Response(data, 200, "success get data", nil)
}

func (svc *vehicles_service) SortType(typeVehicle string) *libs.Resp {
	data, err := svc.repo.Type(typeVehicle)

	if err != nil {
		return libs.Response(data, 400, "failed get data", err)
	}

	return libs.Response(data, 200, "success get data", nil)
}