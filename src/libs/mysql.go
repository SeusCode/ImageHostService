package libs

import (
	"fmt"
	"log"
	"os"
	"restapi/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	Charset  string
}

func (c *DbConfig) InitMysqlDB() *gorm.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database, c.Charset)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}

	db.AutoMigrate(&models.User{}, &models.Image{}, &models.SystemLog{})

	return db
}
