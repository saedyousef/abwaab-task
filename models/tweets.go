// models/users.go

package models

type Tweet struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Body string `json:"body"`
	UserID uint64 `json:"user_id"`
}