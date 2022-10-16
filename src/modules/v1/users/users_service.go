package users

import (
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"github.com/wildanfaz/vehicle_rental/src/interfaces"
	"github.com/wildanfaz/vehicle_rental/src/libs"
)

type users_service struct {
	repo interfaces.UsersRepo
}

func NewService(repo interfaces.UsersRepo) *users_service {
	return &users_service{repo}
}

func (svc *users_service) GetAllUsers() *libs.Resp {
	data, err := svc.repo.FindAllUsers()

	if err != nil {
		return libs.Response(data, 400, "failed get data", err)
	}

	return libs.Response(data, 200, "success get data", nil)
}

func (svc *users_service) AddUser(body *models.User) *libs.Resp {
	if body.Name == "" || body.Email == "" || body.Password == "" {
		return libs.Response(nil, 400, "data can't empty", nil)
	}

	hashpassword, errhash := libs.HashingPassword(body.Password)

	if errhash != nil {
		return libs.Response(nil, 400, "failed hash password", errhash)
	}

	body.Password = hashpassword

	if check := svc.repo.CheckDB(body); check != nil {
		return libs.Response(nil, 400, "failed add data", check)
	}

	result, err := svc.repo.SaveUser(body)

	if err != nil {
		return libs.Response(nil, 400, "failed add data", err)
	}

	return libs.Response(result.Name, 201, "success add data", nil)
}

func (svc *users_service) UpdateUser(vars string, body *models.User) *libs.Resp {
	if checkVars := svc.repo.CheckVars(vars); checkVars != nil {
		return libs.Response(nil, 400, "failed update data", checkVars)
	}

	if check := svc.repo.CheckDB(body); check != nil {
		return libs.Response(nil, 400, "failed update data", check)
	}

	_, err := svc.repo.ChangeUser(vars, body)

	if err != nil {
		return libs.Response(nil, 400, "failed update data", err)
	}

	return libs.Response(nil, 200, "success update data", nil)
}

func (svc *users_service) DeleteUser(vars string, body *models.User) *libs.Resp {
	if check := svc.repo.CheckNameDB(vars); check != nil {
		return libs.Response(nil, 400, "failed delete data", check)
	}

	_, err := svc.repo.RemoveUser(vars, body)

	if err != nil {
		return libs.Response(nil, 400, "failed delete data", err)
	}

	return libs.Response(nil, 200, "success delete data", nil)
}

func (svc *users_service) GetUserByName(name string) *libs.Resp {
	data, err := svc.repo.FindUserByName(name)

	if err != nil {
		return libs.Response(nil, 404, "failed get data", err)
	}

	return libs.Response(data, 200, "success get data", nil)
}

func (svc *users_service) GetUserByEmail(email string) *libs.Resp {
	data, err := svc.repo.FindUserByEmail(email)

	if err != nil {
		return libs.Response(nil, 404, "failed get data", err)
	}

	return libs.Response(data, 200, "success get data", nil)
}

// func (re *users_service) SearchUser(r *http.Request) (*models.Users, error) {
// 	data, err := re.repo.FindUser(r)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }
