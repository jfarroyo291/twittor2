package middlew

import (
	"net/http"

	"github.com/jfarroyo291/twittor2/bd"
)

/*ChequeoBD es el middlew que me permite concer el estada de la BD*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
