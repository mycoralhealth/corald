package web

import (
	"database/sql"
	"net/http"

	"github.com/mycoralhealth/corald/auth0"
)

func handleSession(w http.ResponseWriter, r *http.Request, dbCon *sql.DB) {
	access_token := r.Header.Get("X-Mycoral-Accesstoken")

	user_info, err := auth0.Validate(access_token)
	if err == auth0.Unauthorized {
		handleError(w, r, http.StatusUnauthorized, "")
		return
	} else if err != nil {
		handleError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// TODO: Lookup corresponding user in our database
	// Create entry if we don't have one

	respondWithJSON(w, r, http.StatusOK, user_info)
}
