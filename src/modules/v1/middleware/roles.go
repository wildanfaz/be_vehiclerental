package middleware

// func UserAdmin(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		roles := context.WithValue(r.Context(), "userAdmin", "User Admin")

// 		next.ServeHTTP(w, r.WithContext(roles))
// 	}
// }

// func User(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		roles := context.WithValue(r.Context(), "userAdmin", "User")

// 		next.ServeHTTP(w, r.WithContext(roles))
// 	}
// }

// func Admin(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		roles := context.WithValue(r.Context(), "userAdmin", "Admin")

// 		next.ServeHTTP(w, r.WithContext(roles))
// 	}
// }
