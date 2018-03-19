package auth0

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// UserInfo stores the response from Auth0's userinfo endpoint
// https://auth0.com/docs/api/authentication#get-user-info
type UserInfo struct {
	ClientID      string    `json:"client_id"`
	CreatedAt     time.Time `json:"created_at"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	Identities    []struct {
		Connection string `json:"connection"`
		IsSocial   bool   `json:"isSocial"`
		Provider   string `json:"provider"`
		UserID     string `json:"user_id"`
	} `json:"identities"`
	Name      string    `json:"name"`
	Nickname  string    `json:"nickname"`
	Picture   string    `json:"picture"`
	Sub       string    `json:"sub"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id"`

	AppMetadata struct {
		Admin bool `json:"admin"`
	} `json:"app_metadata"`

	UserMetadata map[string]interface{} `json:"user_metadata"`
}

// IsAdmin returns true if the account has the Admin flag set
// to `true` in the Auth0 app metadata.
func (p UserInfo) IsAdmin() bool {
	return p.AppMetadata.Admin
}

// Unauthorized is an error value to represent 401 Unauthorized response from Auth0
var Unauthorized error

func init() {
	Unauthorized = fmt.Errorf("Unauthorized")
}

// Validate validates the given accessToken by making a request to Auth0.
// If it is invalid, it returns Unauthorized error.
// If it is valid, it returns the user's UserInfo.
func Validate(accessToken string) (UserInfo, error) {
	url := os.Getenv("CORALD_AUTH0_DOMAIN") + "/userinfo?access_token=" + accessToken
	resp, err := http.Get(url)
	if err != nil {
		return UserInfo{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return UserInfo{}, Unauthorized
	} else if resp.StatusCode != 200 {
		return UserInfo{}, fmt.Errorf("Unexpected Auth0 Response: %d", resp.StatusCode)
	}

	var u UserInfo
	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return UserInfo{}, err
	}

	return u, nil
}
