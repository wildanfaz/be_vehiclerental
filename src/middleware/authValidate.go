package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/wildanfaz/vehicle_rental/src/libs"
)

func CheckAuth(roles ...string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			headerToken := r.Header.Get("Authorization")

			if !strings.Contains(headerToken, "Bearer") {
				libs.Response(nil, 401, "invalid header", nil).Send(w)
				return
			}

			token := strings.ReplaceAll(headerToken, "Bearer ", "")

			checkToken, err := libs.CheckToken(token)

			if err != nil {
				libs.Response(nil, 401, "invalid token", err).Send(w)
				return
			}

			// roles := fmt.Sprint(res)
			// range strings.Split(roles, " ")
			var checkRole bool
			for _, v := range roles {
				if strings.ToLower(v) == strings.ToLower(checkToken.Role) {
					checkRole = true
					break
				}
			}

			if !checkRole {
				libs.Response(nil, 401, "unauthorized role", nil).Send(w)
				return
			}

			ctx := context.WithValue(r.Context(), "email", checkToken.Email)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}

// func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		headerToken := r.Header.Get("Authorization")

// 		if !strings.Contains(headerToken, "Bearer") {
// 			libs.Response(nil, 401, "invalid header", nil).Send(w)
// 			return
// 		}

// 		token := strings.ReplaceAll(headerToken, "Bearer ", "")

// 		checkToken, err := libs.CheckToken(token)

// 		if err != nil {
// 			libs.Response(nil, 401, "invalid token", err).Send(w)
// 			return
// 		}

// 		res := Roles(r)
// 		roles := fmt.Sprint(res)
// 		var checkRole bool
// 		for _, v := range strings.Split(roles, " ") {
// 			if strings.ToLower(v) == strings.ToLower(checkToken.Role) {
// 				checkRole = true
// 				break
// 			}
// 		}

// 		if !checkRole {
// 			libs.Response(nil, 401, "unauthorized role", nil).Send(w)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), "name", checkToken.Name)

// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	}
// }

// func Roles(r *http.Request) interface{} {
// 	if roles := r.Context().Value("userAdmin"); roles != nil {
// 		return roles
// 	} else if roles := r.Context().Value("user"); roles != nil {
// 		return roles
// 	} else if roles := r.Context().Value("admin"); roles != nil {
// 		return roles
// 	} else {
// 		return nil
// 	}
// }
