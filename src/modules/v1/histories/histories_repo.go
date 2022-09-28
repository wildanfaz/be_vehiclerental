package histories

import (
	"errors"

	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"gorm.io/gorm"
)

type histories_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *histories_repo {
	return &histories_repo{db}
}

func (re *histories_repo) FindAllHistories() (*models.Histories, error) {
	var data models.Histories

	result := re.db.Order("created_at desc").Preload("Vehicle", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("vehicle_id, vehicle_name, created_at, updated_at, total_rented")
	}).Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("user_id, name, email, created_at, updated_at")
	}).Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get users")
	}

	return &data, nil
}

func (re *histories_repo) SaveHistory(body *models.History) (*models.History, error) {
	var vehicle models.Vehicle

	result := re.db.Create(body)

	if result.Error != nil {
		return nil, errors.New("failed save history")
	}

	re.db.Where("vehicle_id = ?", body.VehicleId).First(&vehicle)

	re.db.Model(&vehicle).Where("vehicle_id = ?", body.VehicleId).Update("total_rented", vehicle.TotalRented+1)

	return body, nil
}

func (re *histories_repo) ChangeHistory(vars string, body *models.History) (*models.History, error) {
	var check int64

	re.db.Model(&body).Where("history_id = ?", vars).Count(&check)
	checkName := check > 0

	if checkName == false {
		return nil, errors.New("history is not exists")
	}

	result := re.db.Model(&body).Where("history_id = ?", vars).Updates(body)

	if result.Error != nil {
		return nil, errors.New("failed update history")
	}

	return body, nil
}

func (re *histories_repo) RemoveHistory(vars string, body *models.History) (*models.History, error) {
	var vehicle models.Vehicle

	var check int64

	re.db.Model(&body).Where("history_id = ?", vars).Count(&check)
	checkName := check > 0

	if checkName == false {
		return nil, errors.New("history is not exists")
	}

	re.db.Where("history_id = ?", vars).First(&body)

	re.db.Where("vehicle_id = ?", body.VehicleId).First(&vehicle)

	re.db.Model(&vehicle).Where("vehicle_id = ?", body.VehicleId).Update("total_rented", vehicle.TotalRented-1)

	result := re.db.Delete(body, vars)

	if result.Error != nil {
		return nil, errors.New("failed delete history")
	}

	return body, nil
}

func (re *histories_repo) FindHistory(search string) (*models.Histories, error) {
	var data models.Histories

	result := re.db.Preload("Vehicle", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("vehicle_id, vehicle_name, created_at, updated_at, total_rented")
	}).Preload("User", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("user_id, name, email, created_at, updated_at")
	}).Where("vehicle_id = ?", search).Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get users")
	}

	return &data, nil
}
