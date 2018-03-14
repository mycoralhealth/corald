package web

import (
	"io"
	"net/http"
	"net/url"
	"os"
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

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func generateProxiedURL(u *url.URL) *url.URL {
	u.Scheme = "http"
	u.Host = os.Getenv("CORALD_IPFS_API_HOSTNAME") + ":" + os.Getenv("CORALD_IPFS_API_PORT")
	return u
}

func handleIPFSAdd(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	if !addThrottle.Bump(u.Name) {
		handleError(w, r, http.StatusTooManyRequests, "You have exceeded your limit for today")
		return
	}

	// Make a POST request to the IPFS server, including the body that we received
	url := generateProxiedURL(r.URL).String()
	resp, err := http.Post(url, "text/plain", r.Body)
	if err != nil {
		handleError(w, r, http.StatusServiceUnavailable, err.Error())
		return
	}
	defer resp.Body.Close()

	// Copy the response from the IPFS server back to our client
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, r.Body)
}

func handleIPFSCat(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	// Make a GET request to the IPFS Server.
	url := generateProxiedURL(r.URL).String()
	resp, err := http.Get(url)
	if err != nil {
		handleError(w, r, http.StatusServiceUnavailable, err.Error())
		return
	}
	defer resp.Body.Close()

	// Copy the response from the IPFS server back to our client
	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, r.Body)
}
