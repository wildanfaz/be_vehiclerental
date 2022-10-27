package interfaces

import (
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"github.com/wildanfaz/vehicle_rental/src/libs"
)

type UsersRepo interface {
	FindAllUsers() (*models.Users, error)
	FindUserByName(name string) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
	FindUserByEmail2(email string) (*models.User, error)
	SaveUser(body *models.User) (*models.User, error)
	ChangeUser(vars string, body *models.User) (*models.User, error)
	RemoveUser(vars string, body *models.User) (*models.User, error)
	CheckDB(body *models.User) error
	CheckVars(vars string) error
	CheckNameDB(vars string) error
	CheckEmailDB(vars string) error
}

type UsersService interface {
	GetAllUsers() *libs.Resp
	GetUserByName(name string) *libs.Resp
	GetUserByEmail(email string) *libs.Resp
	AddUser(body *models.User) *libs.Resp
	UpdateUser(vars string, body *models.User) *libs.Resp
	DeleteUser(vars string, body *models.User) *libs.Resp
}
