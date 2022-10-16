package vehicles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"github.com/wildanfaz/vehicle_rental/src/interfaces"
	"github.com/wildanfaz/vehicle_rental/src/libs"
)

type vehicles_ctrl struct {
	svc interfaces.VehiclesService
}

func NewCtrl(svc interfaces.VehiclesService) *vehicles_ctrl {
	return &vehicles_ctrl{svc}
}

func (ctrl *vehicles_ctrl) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	data := ctrl.svc.GetAllVehicles()

	if data.Error != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *vehicles_ctrl) AddVehicle(w http.ResponseWriter, r *http.Request) {
	var datas models.Vehicle

	imageName := r.Context().Value("imageName")

	datas.Image = fmt.Sprint(imageName)

	// r.ParseMultipartForm(20 << 20)

	//**decode from multipart/form-data
	if err := schema.NewDecoder().Decode(&datas, r.MultipartForm.Value); err != nil {
		libs.Response(nil, 400, "failed to decode", err).Send(w)
		return
	}

	data := ctrl.svc.AddVehicle(&datas)

	if data.Error != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *vehicles_ctrl) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	var datas models.Vehicle

	bd, errs := ioutil.ReadAll(r.Body)
	if errs != nil {
		libs.Response(nil, 400, "failed to readAll", errs).Send(w)
	}

	err := json.Unmarshal(bd, &datas)
	if err != nil {
		libs.Response(nil, 400, "failed to Unmarshal", err).Send(w)
	}

	// err := json.NewDecoder(r.Body).Decode(&datas)
	// if err != nil {
	// 	libs.Response(nil, 400, "failed to decode", err).Send(w)
	// 	return
	// }
	
	vars := mux.Vars(r)
	data := ctrl.svc.UpdateVehicle(vars["vehicle_id"], &datas)

	if data.Error != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *vehicles_ctrl) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	var datas models.Vehicle

	vars := mux.Vars(r)
	data := ctrl.svc.DeleteVehicle(vars["vehicle_id"], &datas)

	if data.Error != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *vehicles_ctrl) SearchVehicle(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("vehicle_name")
	data := ctrl.svc.SearchVehicle(search)

	if data.Error != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *vehicles_ctrl) PopularVehicles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := vars["offset"]
	offset,_ := strconv.Atoi(res)
	data := ctrl.svc.PopularVehicles(offset)

	if data.Error != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *vehicles_ctrl) GetVehicleDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	data := ctrl.svc.GetVehicleDetail(id)

	if data.Error != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *vehicles_ctrl) TypeVehicles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeVehicle := vars["type"]
	data := ctrl.svc.TypeVehicles(typeVehicle)

	if data.Error != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *vehicles_ctrl) SortLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	location := vars["location"]
	data := ctrl.svc.SortLocation(location)

	if data.Error != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *vehicles_ctrl) SortType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	typeVehicle := vars["type"]
	data := ctrl.svc.SortType(typeVehicle)

	if data.Error != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}