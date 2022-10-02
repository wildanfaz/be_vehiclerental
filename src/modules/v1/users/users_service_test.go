package users

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
)

func TestGetAllUsers(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Users{
		{Name: "user", Email: "user@gmail.com"},
		{Name: "admin", Email: "admin@gmail.com"},
	}

	repo.mock.On("FindAllUsers").Return(&dataMock, nil)

	data := service.GetAllUsers()

	fmt.Println(data)
	result := data.Data.(*models.Users)

	for i, v := range *result {
		assert.Equal(t, dataMock[i].Name, v.Name, "name is different")
	}
}

func TestGetUserByName(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{
		Name:  "user",
		Email: "user@gmail.com",
	}

	repo.mock.On("FindUserByName", "user").Return(&dataMock, nil)

	data := service.GetUserByName("user")

	fmt.Println(data)
	result := data.Data.(*models.User)
	assert.Equal(t, "user", result.Name, "name is different")
}

func TestAddUser(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{
		Name:  "faz",
		Email: "faz@gmail.com",
	}

	repo.mock.On("SaveUser", &dataMock).Return(&dataMock, nil)
	data := service.AddUser(&dataMock)

	var expectName = "faz"
	result := data.Data.(string)

	assert.Equal(t, expectName, result, "name is not faz")
	assert.IsType(t, "string", result, "result is not string")

	// typeInt := 123
	// assert.IsType(t, typeInt, "num", "actual is not int")
}

func TestUpdateUser(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{
		Name:  "faz1",
		Email: "faz1@gmail.com",
	}

	repo.mock.On("ChangeUser", "faz", &dataMock).Return(&dataMock, nil)
	data := service.UpdateUser("faz", &dataMock)

	assert.Equal(t, 200, data.Status, "status is not 200")
}

func TestDeleteUser(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{
		Name:  "faz1",
		Email: "faz1@gmail.com",
	}

	repo.mock.On("RemoveUser", "faz1", &dataMock).Return(&dataMock, nil)
	data := service.DeleteUser("faz1", &dataMock)

	assert.Equal(t, 200, data.Status, "status is not 200")
}
