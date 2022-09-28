package histories

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wildanfaz/vehicle_rental/src/database/orm/models"
	"github.com/wildanfaz/vehicle_rental/src/interfaces"
	"github.com/wildanfaz/vehicle_rental/src/libs"
)

type histories_ctrl struct {
	svc interfaces.HistoriesService
}

func NewCtrl(svc interfaces.HistoriesService) *histories_ctrl {
	return &histories_ctrl{svc}
}

func (ctrl *histories_ctrl) GetAllHistories(w http.ResponseWriter, r *http.Request) {
	data := ctrl.svc.GetAllHistories()

	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *histories_ctrl) AddHistory(w http.ResponseWriter, r *http.Request) {
	var datas models.History

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Response(nil, 400, "failed to decode", err).Send(w)
		return
	}

	data := ctrl.svc.AddHistory(&datas)

	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *histories_ctrl) UpdateHistory(w http.ResponseWriter, r *http.Request) {
	var datas models.History

	err := json.NewDecoder(r.Body).Decode(&datas)
	if err != nil {
		libs.Response(nil, 400, "failed to decode", err).Send(w)
		return
	}
	vars := mux.Vars(r)
	data := ctrl.svc.UpdateHistory(vars["history_id"], &datas)

	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *histories_ctrl) DeleteHistory(w http.ResponseWriter, r *http.Request) {
	var datas models.History

	vars := mux.Vars(r)
	data := ctrl.svc.DeleteHistory(vars["history_id"], &datas)

	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}

func (ctrl *histories_ctrl) SearchHistory(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("vehicle_id")
	data := ctrl.svc.SearchHistory(search)

	if data.IsError != nil {
		data.Send(w)
		return
	}

	data.Send(w)
}
