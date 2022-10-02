package vehicles

import (
	"errors"

	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"gorm.io/gorm"
)

type vehicles_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *vehicles_repo {
	return &vehicles_repo{db}
}

func (re *vehicles_repo) FindAllVehicles() (*models.Vehicles, error) {
	var data models.Vehicles

	result := re.db.Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get vehicles")
	}

	return &data, nil
}

func (re *vehicles_repo) SaveVehicle(body *models.Vehicle) (*models.Vehicle, error) {
	result := re.db.Create(body)

	if result.Error != nil {
		return nil, errors.New("failed save vehicle")
	}

	return body, nil
}

func (re *vehicles_repo) ChangeVehicle(vars string, body *models.Vehicle) (*models.Vehicle, error) {
	// var check int64

	// re.db.Model(&body).Where("vehicle_id = ?", vars).Count(&check)
	// checkName := check > 0

	// if checkName == false {
	// 	return nil, errors.New("vehicle is not exists")
	// }

	result := re.db.Model(&body).Where("vehicle_id = ?", vars).Updates(body)

	if result.Error != nil {
		return nil, errors.New("failed update vehicle")
	}

	return body, nil
}

func (re *vehicles_repo) RemoveVehicle(vars string, body *models.Vehicle) (*models.Vehicle, error) {
	// var check int64

	// re.db.Model(&body).Where("vehicle_id = ?", vars).Count(&check)
	// checkName := check > 0

	// if checkName == false {
	// 	return nil, errors.New("vehicle is not exists")
	// }

	result := re.db.Delete(body, vars)

	if result.Error != nil {
		return nil, errors.New("failed delete vehicle")
	}

	return body, nil
}

func (re *vehicles_repo) FindVehicle(search string) (*models.Vehicles, error) {
	var data models.Vehicles

	s := "%" + search + "%"
	result := re.db.Where("LOWER(vehicle_name) LIKE ?", s).Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed search vehicles")
	}

	return &data, nil
}

func (re *vehicles_repo) RatingVehicles() (*models.Vehicles, error) {
	var data models.Vehicles

	result := re.db.Order("rating desc, total_rented desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get vehicles")
	}

	return &data, nil
}

func (re *vehicles_repo) CheckId(vars string, body *models.Vehicle) error {
	var check int64

	re.db.Model(&body).Where("vehicle_id = ?", vars).Count(&check)
	checkName := check > 0

	if checkName == false {
		return errors.New("vehicle is not exists")
	}

	return nil
}
