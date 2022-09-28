package middleware

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/wildanfaz/vehicle_rental/src/libs"
)

// **img
func UploadFileImage(next http.HandlerFunc) http.HandlerFunc {
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

		//**check
		// handlerFile.Header
		fmt.Println(handlerFile.Header.Get("Content-Type"))

		name := strings.ReplaceAll(strings.ReplaceAll(time.Now().Format(time.ANSIC), ":", "-")+"-"+handlerFile.Filename, " ", "_")
		result, errs := os.Create("images/" + name)

		if errs != nil {
			libs.Response(nil, 400, "failed create file", errs).Send(w)
			return
		}

		if _, error := io.Copy(result, file); error != nil {
			libs.Response(nil, 400, "failed copy file", errs).Send(w)
			return
		}

		libs.Response(nil, 200, "success upload file", nil).Send(w)

		ctx := context.WithValue(r.Context(), "imageName", name)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
