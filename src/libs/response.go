package libs

import (
	"encoding/json"
	"net/http"
)

type Resp struct {
	Status      int         `json:"status"`
	Description string      `json:"description"`
	Message     string      `json:"message"`
	IsError     interface{} `json:"error,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

func (res *Resp) Send(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(res.Status)

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("error encode"))
	}
}

func Response(data interface{}, status int, message string, isError error) *Resp {
	if isError != nil {
		return &Resp{
			Status:      status,
			Description: statusDescription(status),
			Message:     message,
			IsError:     isError.Error(),
			Data:        data,
		}
	}

	return &Resp{
		Status:      status,
		Description: statusDescription(status),
		Message:     message,
		IsError:     isError,
		Data:        data,
	}
}

func statusDescription(status int) string {
	switch status {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 304:
		return "Not Modified"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	case 501:
		return "Bad Gateway"
	default:
		return ""
	}
}
