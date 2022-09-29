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
		assert.Equal(t, dataMock[i].Name, v.Name, "name is not match")
	}
}

func TestGetUserByName(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{
		Name:  "user",
		Email: "user@gmail.com",
	}

	repo.mock.On("FindUserByName").Return(&dataMock, nil)

	data := service.GetUserByName("user")

	fmt.Println(data)
	// result := data.Data.(*models.Users)
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
