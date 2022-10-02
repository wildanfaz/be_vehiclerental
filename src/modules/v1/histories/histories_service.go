package histories

import (
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"github.com/wildanfaz/vehicle_rental/src/interfaces"
	"github.com/wildanfaz/vehicle_rental/src/libs"
)

type histories_service struct {
	repo interfaces.HistoriesRepo
}

func NewService(repo interfaces.HistoriesRepo) *histories_service {
	return &histories_service{repo}
}

func (svc *histories_service) GetAllHistories() *libs.Resp {
	data, err := svc.repo.FindAllHistories()

	if err != nil {
		return libs.Response(data, 400, "failed get data", err)
	}

	return libs.Response(data, 200, "success get data", nil)
}

func (svc *histories_service) AddHistory(data *models.History) *libs.Resp {
	_, err := svc.repo.SaveHistory(data)

	if err != nil {
		return libs.Response(nil, 400, "failed add data", err)
	}

	return libs.Response(nil, 200, "success add data", nil)
}

func (svc *histories_service) UpdateHistory(vars string, data *models.History) *libs.Resp {
	if check := svc.repo.CheckId(vars, data); check != nil {
		return libs.Response(nil, 400, "failed update data", check)
	}

	_, err := svc.repo.ChangeHistory(vars, data)

	if err != nil {
		return libs.Response(nil, 400, "failed update data", err)
	}

	return libs.Response(nil, 201, "success update data", nil)
}

func (svc *histories_service) DeleteHistory(vars string, data *models.History) *libs.Resp {
	_, err := svc.repo.RemoveHistory(vars, data)

	if err != nil {
		return libs.Response(nil, 400, "failed delete data", err)
	}

	return libs.Response(nil, 200, "success delete data", nil)
}

func (svc *histories_service) SearchHistory(search string) *libs.Resp {
	data, err := svc.repo.FindHistory(search)

	if err != nil {
		return libs.Response(data, 404, "failed search data", err)
	}

	return libs.Response(data, 200, "success search data", nil)
}
