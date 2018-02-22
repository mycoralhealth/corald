package model

// User is a mocked Ethereum account representing a patient
// in the mycoralhealth system
type User struct {
	Username  string // NOT NULL with no default
	Email     string
	Address   string
	PublicKey string
}
