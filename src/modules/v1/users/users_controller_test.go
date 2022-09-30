package users

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"github.com/wildanfaz/vehicle_rental/src/libs"
)

var repos = RepoMock{mock.Mock{}}
var service = NewService(&repos)
var ctrl = NewCtrl(service)

var dataMock = models.User{
	Name:  "user",
	Email: "user@gmail.com",
}

func TestCtrlGetUserByName(t *testing.T) {
	w := httptest.NewRecorder()
	mux := mux.NewRouter()

	repos.mock.On("FindUserByName", "user").Return(&dataMock, nil)

	mux.HandleFunc("/test/users", ctrl.GetUserByName)

	mux.ServeHTTP(w, httptest.NewRequest("GET", "/test/users", nil))

	var prods models.User
	response := libs.Resp{
		Data: &prods,
	}

	fmt.Println(response)

	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	//**check
	// response.Error = "err"
	assert.Equal(t, 200, w.Code, "status is not 200")
	assert.Nil(t, response.Error, "error is not nil")
}
