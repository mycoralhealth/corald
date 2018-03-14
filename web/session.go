package web

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mycoralhealth/corald/auth0"
	"github.com/mycoralhealth/corald/model"
)

func handleSession(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {

	// Create entry if we don't have one
	var user model.User
	dbCon.Where(model.User{Name: u.Name}).Assign(model.User{LastLogin: time.Now()}).FirstOrCreate(&user)

	respondWithJSON(w, r, http.StatusOK, u)
}
