package models

import "time"

type User struct {
	Id       uint      `json:"id" gorm:"primary_key"`
	Created  time.Time `json:"created" gorm:"autoCreateTime"`
	Updated  time.Time `json:"updated" gorm:"autoUpdateTime"`
	Name     string    `json:"name" gorm:"type:varchar(32);not null"`
	PassHash string    `json:"pass" gorm:"type:varchar(32);not null"`
}
