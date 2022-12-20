package models

import "time"

type Image struct {
	//ID autoincremental
	Id uint `json:"id" gorm:"primary_key;not null"`
	//Upload UID
	Uid string `json:"uid" gorm:"type:varchar(255);not null"`
	//Insertion date
	Created time.Time `json:"created" gorm:"autoCreateTime;not null"`
	//Update date
	Updated time.Time `json:"updated" gorm:"autoUpdateTime;not null"`
	//Transfer counter date (reset every X minutes... ¿5?)
	TransferDate time.Time `json:"TransferDate" gorm:"not null"`
	//Transfer bytes
	TransferSize uint `json:"TransferSize" gorm:"not null"`
	//Transfer bytes
	TotalTransferSize uint `json:"TotalTransferSize" gorm:"not null"`
	//Image has thumbnail? (%uid% + "_thumb_" + %filename% )
	Thumb bool `json:"thumb" gorm:"not null"`
	//Filename
	Name string `json:"name" gorm:"type:varchar(255);not null"`
	//Description (optional)
	Desc string `json:"desc" gorm:"type:text;not null"`
}
