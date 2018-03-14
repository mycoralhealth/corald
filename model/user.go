package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// User is a mocked Ethereum account representing a patient
// in the mycoralhealth system
type User struct {
	gorm.Model
	Name      string
	Email     string
	Address   string
	PublicKey string
	LastLogin time.Time
}
