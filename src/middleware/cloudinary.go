package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
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

		name := strings.ReplaceAll(strings.ReplaceAll(time.Now().Format(time.ANSIC), ":", "-")+"-"+handlerFile.Filename, " ", "_")

		cld, errs := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_CLOUD_NAME"),os.Getenv("CLOUDINARY_API_KEY"),os.Getenv("CLOUDINARY_API_SECRET"))

		if errs != nil {
			libs.Response(nil, 400, "err env cloudinary", errs).Send(w)
			return
		}

		upload, err := cld.Upload.Upload(nil, file, uploader.UploadParams{Folder: "vehiclerental", PublicID: name})

		libs.Response(nil, 200, "success upload file", nil).Send(w)

		ctx := context.WithValue(r.Context(), "imageName", upload.SecureURL)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}