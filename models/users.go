package models

import "time"

type Users struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username"`
	Password  string    `json:"password" gorm:"encrypted"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
