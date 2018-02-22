package db

import (
	"database/sql"

	"github.com/mycoralhealth/mycoral-patient-server/model"
)

// GetConstituent returns info for a single canstituent
func GetUser(db *sql.DB, username string) (model.User, error) {

	row := db.QueryRow(`SELECT
		username, email, address, public_key
	FROM users WHERE username = $1
	`, username)
	user, err := scanUser(row)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func scanUser(row rowScanner) (model.User, error) {
	var user model.User
	if err := row.Scan(
		&user.Username, &user.Email, &user.Address, &user.PublicKey,
	); err != nil {
		return user, err
	}

	return user, nil
}
