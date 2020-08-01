// models/users.go

package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

