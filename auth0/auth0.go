package auth0

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type UserInfo struct {
	ClientID      string
	CreatedAt     time.Time `json:"created_at"`
	Email         string
	EmailVerified bool `json:"email_verified"`
	Identities    []struct {
		Connection string
		IsSocial   bool
		Provider   string
		UserId     string `json:"user_id"`
	}
	Name      string
	Nickname  string
	Picture   string
	Sub       string
	UpdatedAt time.Time `json:"updated_at"`
	UserId    string    `json:"user_id"`

	AppMetadata struct {
		Admin bool
	} `json:"app_metadata"`

	UserMetadata map[string]interface{} `json:"user_metadata"`
}

func (p UserInfo) IsAdmin() bool {
	return p.AppMetadata.Admin
}

var Unauthorized error

func init() {
	Unauthorized = fmt.Errorf("Unauthorized")
}

func Validate(access_token string) (UserInfo, error) {
	url := os.Getenv("CORALD_AUTH0_DOMAIN") + "/userinfo?access_token=" + access_token
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
