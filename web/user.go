package web

import (
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/mycoralhealth/corald/auth0"
	"github.com/mycoralhealth/corald/model"

	"github.com/gorilla/mux"
)

func handleGetAllUsers(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	handleError(w, r, http.StatusNotImplemented, "")
}

func handleGetUser(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	//FIXME: This should not be a public endpoint

	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		handleError(w, r, http.StatusNotFound, err.Error())
		return
	}

	var user model.User
	err = dbCon.First(&user, userid).Error
	if gorm.IsRecordNotFoundError(err) {
		handleError(w, r, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		handleError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, r, http.StatusOK, user)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	handleError(w, r, http.StatusNotImplemented, "")
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	handleError(w, r, http.StatusNotImplemented, "")
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	handleError(w, r, http.StatusNotImplemented, "")
}
