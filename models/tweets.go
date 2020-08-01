// models/users.go

package models

import (
	"github.com/jinzhu/gorm"
)

type Tweet struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Body string `json:"body"`
}
