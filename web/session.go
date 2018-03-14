package web

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mycoralhealth/corald/auth0"
	"github.com/mycoralhealth/corald/model"
)

func handleSession(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB) {
	accessToken := r.Header.Get("X-Mycoral-Accesstoken")

	userInfo, err := auth0.Validate(accessToken)
	if err == auth0.Unauthorized {
		handleError(w, r, http.StatusUnauthorized, "")
		return
	} else if err != nil {
		handleError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Create entry if we don't have one
	var user model.User
	dbCon.Where(model.User{Name: userInfo.Name}).Assign(model.User{LastLogin: time.Now()}).FirstOrCreate(&user)

	respondWithJSON(w, r, http.StatusOK, userInfo)
}
