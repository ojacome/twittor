package middlewares

import (
	"github.com/ojacome/twittor/db"
	"net/http"
)

// CheckBD middleware que permite conoceer el estado de la base.
func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "error en la conexión a mongoDB", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
