package web

import (
	"database/sql"
	"net/http"

	"github.com/mycoralhealth/corald/db"

	"github.com/gorilla/mux"
)

func handleGetAllUsers(w http.ResponseWriter, r *http.Request, dbCon *sql.DB) {
	handleError(w, r, http.StatusNotImplemented, "")
}

func handleGetUser(w http.ResponseWriter, r *http.Request, dbCon *sql.DB) {
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := db.GetUser(dbCon, username)
	if err == sql.ErrNoRows {
		handleError(w, r, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		handleError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, r, http.StatusOK, user)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request, dbCon *sql.DB) {
	handleError(w, r, http.StatusNotImplemented, "")
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request, dbCon *sql.DB) {
	handleError(w, r, http.StatusNotImplemented, "")
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request, dbCon *sql.DB) {
	handleError(w, r, http.StatusNotImplemented, "")
}
