package auth0

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type UserInfo struct {
	Sub       string `json:"sub"`
	Nickname  string `json:"nickname"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	UpdatedAt string `json:"updated_at"`
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
