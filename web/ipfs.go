package web

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
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

func generateProxiedURL(u *url.URL) *url.URL {
	u.Scheme = "http"
	u.Host = os.Getenv("CORALD_IPFS_API_HOSTNAME") + ":" + os.Getenv("CORALD_IPFS_API_PORT")
	u.Path = "/api/v0/" + u.Path[9:]
	return u
}

func generateProxiedRequest(inReq *http.Request) (*http.Request, error) {
	// Unpack file from incoming request
	inFile, inHeader, err := inReq.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	log.Printf("IPFS add: %s\n", inHeader.Filename)

	// Encode file as multipart
	var b bytes.Buffer
	mw := multipart.NewWriter(&b)
	fw, err := mw.CreateFormFile("file", inHeader.Filename)
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, inFile); err != nil {
		return nil, err
	}
	if err = mw.Close(); err != nil {
		return nil, err
	}

	// Create an outgoing request
	url := generateProxiedURL(inReq.URL).String()
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())

	return req, nil
}

func handleIPFSAdd(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	if !addThrottle.Bump(u.Name) {
		handleError(w, r, http.StatusTooManyRequests, "You have exceeded your limit for today")
		return
	}

	req, err := generateProxiedRequest(r)
	if err != nil {
		handleError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	// Submit the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		handleError(w, r, http.StatusServiceUnavailable, err.Error())
		return
	}
	defer resp.Body.Close()

	// If the IPFS server returned an error, log it
	// and return a generic error to our client
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		bod, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			handleError(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		handleError(w, r, http.StatusInternalServerError, string(bod))
		return
	}

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		handleError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}

func handleIPFSCat(w http.ResponseWriter, r *http.Request, dbCon *gorm.DB, u *auth0.UserInfo) {
	log.Printf("IPFS cat: %s\n", r.URL)

	// Make a GET request to the IPFS Server.
	url := generateProxiedURL(r.URL).String()
	resp, err := http.Get(url)
	if err != nil {
		handleError(w, r, http.StatusServiceUnavailable, err.Error())
		return
	}
	defer resp.Body.Close()

	// If the IPFS server returned an error, log it
	// and return a generic error to our client
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		bod, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			handleError(w, r, http.StatusInternalServerError, err.Error())
			return
		}
		handleError(w, r, http.StatusInternalServerError, string(bod))
		return
	}

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		handleError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
}
