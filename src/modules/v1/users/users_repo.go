package users

import (
	"errors"

	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"gorm.io/gorm"
)

type users_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *users_repo {
	return &users_repo{db}
}

func (re *users_repo) FindAllUsers() (*models.Users, error) {
	var data models.Users

	result := re.db.Order("created_at desc").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed get users")
	}

	return &data, nil
}

func (re *users_repo) SaveUser(body *models.User) (*models.User, error) {
	// var exists bool

	// re.db.Raw("SELECT EXISTS(SELECT * FROM users WHERE name = ?)", data.Name).Scan(&exists)

	// if exists {
	// 	return nil, errors.New("name already exists")
	// }
	// var exists int64

	// re.db.Model(&body).Where("name = ? OR email = ?", body.Name, body.Email).Count(&exists)
	// isExists := exists > 0

	// if isExists {
	// 	return nil, errors.New("name or email already exists")
	// }

	// hashpassword, err := helpers.Hashing(data.Password)

	result := re.db.Create(body)

	if result.Error != nil {
		return nil, errors.New("failed save data")
	}

	return body, nil
}

func (re *users_repo) ChangeUser(vars string, body *models.User) (*models.User, error) {
	// var exists bool

	// re.db.Raw("SELECT EXISTS (SELECT * FROM users WHERE name = ?)", data.Name).Scan(&exists)

	// if exists == true {
	// 	return nil, errors.New("name already exists")
	// }
	// var check int64

	// re.db.Model(&body).Where("name = ?", vars).Count(&check)
	// checkName := check > 0

	// if checkName == false {
	// 	return nil, errors.New("name is not exists")
	// }

	// var exists int64

	// re.db.Model(&body).Where("name = ? or email = ?", body.Name, body.Email).Count(&exists)
	// isExists := exists > 0

	// if isExists {
	// 	return nil, errors.New("name or email already exists")
	// }

	result := re.db.Model(&body).Where("name = ?", vars).Updates(body)

	if result.Error != nil {
		return nil, errors.New("failed update data")
	}

	return body, nil
}

func (re *users_repo) RemoveUser(vars string, body *models.User) (*models.User, error) {
	// var check int64

	// re.db.Model(&body).Where("name = ?", vars).Count(&check)
	// checkName := check > 0

	// if checkName == false {
	// 	return nil, errors.New("name is not exists")
	// }

	result := re.db.Where("name = ?", vars).Delete(body)

	if result.Error != nil {
		return nil, errors.New("failed delete data")
	}

	return body, nil
}

func (re *users_repo) FindUserByName(name string) (*models.User, error) {
	var user models.User
	var check int64

	re.db.Model(&user).Where("name = ?", name).Count(&check)
	checkName := check > 0

	if checkName == false {
		return nil, errors.New("name is not exists")
	}

	result := re.db.Where("name = ?", name).Find(&user)

	if result.Error != nil {
		return nil, errors.New("name is not exists")
	}

	return &user, nil
}

func (re *users_repo) FindUserByEmail(email string) (*models.User, error) {
	// var check int64

	// re.db.Model(&user).Where("email = ?", email).Count(&check)
	// checkName := check > 0

	// if checkName == false {
	// 	return nil, errors.New("email is not exists")
	// }
	var user models.User

	result := re.db.Where("email = ?", email).Omit("password").Find(&user)

	if result.Error != nil {
		return nil, errors.New("error in repo FindUserByEmail")
	}

	return &user, nil
}

func (re *users_repo) FindUserByEmail2(email string) (*models.User, error) {
	// var check int64

	// re.db.Model(&user).Where("email = ?", email).Count(&check)
	// checkName := check > 0

	// if checkName == false {
	// 	return nil, errors.New("email is not exists")
	// }
	var user models.User

	result := re.db.Where("email = ?", email).Find(&user)

	if result.Error != nil {
		return nil, errors.New("error in repo FindUserByEmail")
	}

	return &user, nil
}

func (re *users_repo) CheckEmailDB(vars string) error {
	var user models.User
	var check int64

	re.db.Model(&user).Where("email = ?", vars).Count(&check)
	checkName := check > 0

	if checkName == false {
		return errors.New("email is not exists")
	}

	return nil
}

func (re *users_repo) CheckNameDB(vars string) error {
	var user models.User
	var check int64

	re.db.Model(&user).Where("name = ?", vars).Count(&check)
	checkName := check > 0

	if checkName == false {
		return errors.New("name is not exists")
	}

	return nil
}

func (re *users_repo) CheckDB(body *models.User) error {
	var exists int64

	re.db.Model(&body).Where("name = ? OR email = ?", body.Name, body.Email).Count(&exists)
	isExists := exists > 0

	if isExists {
		return errors.New("name or email already exists")
	}

	return nil
}

func (re *users_repo) CheckVars(vars string) error {
	var user models.User
	var check int64

	re.db.Model(&user).Where("name = ?", vars).Count(&check)
	checkName := check > 0

	if checkName == false {
		return errors.New("name is not exists")
	}

	return nil
}

// func (re *users_repo) FindUser(r *http.Request) (*models.Users, error) {
// 	var data models.Users

// 	search := r.URL.Query().Get("name")
// 	s := "%" + search + "%"
// 	s = strings.ToLower(s)
// 	result := re.db.Where("LOWER(name) LIKE ?", s).Order("created_at desc").Find(&data)

// 	if result.Error != nil {
// 		return nil, errors.New("failed get users")
// 	}

// 	return &data, nil
// }
