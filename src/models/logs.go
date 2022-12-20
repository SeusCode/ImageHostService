package models

import "time"

type SystemLog struct {
	Id      uint      `json:"id" gorm:"primary_key;not null"`
	Uid     string    `json:"uid" gorm:"type:varchar(255);not null"`
	Created time.Time `json:"created" gorm:"autoCreateTime;not null"`
	Updated time.Time `json:"updated" gorm:"autoUpdateTime;not null"`
	Ip      string    `json:"ip" gorm:"type:varchar(255);not null"`
	Name    string    `json:"name" gorm:"type:varchar(255);not null"`
	Desc    string    `json:"desc" gorm:"type:text;not null"`
}
