package users

import (
	"github.com/stretchr/testify/mock"
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) FindAllUsers() (*models.Users, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Users), nil
}

// **test 1
func (m *RepoMock) FindUserByName(name string) (*models.User, error) {
	if name == "" {
		name = "user"
	}
	args := m.mock.Called(name)
	return args.Get(0).(*models.User), nil
}

// **test 2
func (m *RepoMock) SaveUser(body *models.User) (*models.User, error) {
	args := m.mock.Called(body)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) ChangeUser(vars string, body *models.User) (*models.User, error) {
	args := m.mock.Called(vars, body)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) RemoveUser(vars string, body *models.User) (*models.User, error) {
	args := m.mock.Called(vars, body)
	return args.Get(0).(*models.User), nil
}
