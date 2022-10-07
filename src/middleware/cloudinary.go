package middleware

import (
	"net/http"

	"github.com/wildanfaz/vehicle_rental/src/libs"
)

func CloudinaryAddImg(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//**parse multipart/form-data
		if isErr := r.ParseMultipartForm(20 << 20); isErr != nil {
			libs.Response(nil, 400, "failed parse form", isErr).Send(w)
			return
		}

		file, handlerFile, err := r.FormFile("image")

		defer file.Close()

		if err != nil {
			libs.Response(nil, 400, "failed to upload file", err).Send(w)
			return
		}

		//**file validation
		checkType := handlerFile.Header.Get("Content-Type") == "image/jpeg" || handlerFile.Header.Get("Content-Type") == "image/jpg" || handlerFile.Header.Get("Content-Type") == "image/png"

		if !checkType {
			libs.Response(nil, 400, "invalid file format, only support image file", err).Send(w)
			return
		}
	}
}