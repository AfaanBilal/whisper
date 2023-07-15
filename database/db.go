/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package database

import (
	"os"

	"github.com/AfaanBilal/whisper/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
}

func RunMigrations() {
	DB.AutoMigrate(
		&models.User{},
		&models.AccessToken{},
		&models.Follow{},
		&models.Post{},
		&models.Like{},
		&models.Notification{},
		&models.VerificationCode{},
	)
}
