package web

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mycoralhealth/corald/auth0"
	"github.com/mycoralhealth/corald/utils/throttle"
)

var addThrottle *throttle.Throttle

func init() {
	// Each user can make 10 IPFS add requests per day
	addThrottle = throttle.NewThrottle(50, 24*time.Hour)
}

func handleIPFSAdd(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	if !addThrottle.Bump(u.Name) {
		handleError(w, r, http.StatusTooManyRequests, "You have exceeded your limit for today")
		return
	}

	handleError(w, r, http.StatusNotImplemented, "")
}

func handleIPFSCat(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	handleError(w, r, http.StatusNotImplemented, "")
}
