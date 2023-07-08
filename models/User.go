/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID   uint      `gorm:"primaryKey,autoIncrement"`
	UUID uuid.UUID `gorm:"type:varchar(60);"`

	Name     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"uniqueIndex"`
	Password string `gorm:"type:varchar(255)"`

	Birthday time.Time
	Bio      sql.NullString
	Link     sql.NullString

	Role string `gorm:"type:varchar(255);default:user"`

	ActivatedAt sql.NullTime
	VerifiedAt  sql.NullTime
	Meta        string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.UUID = uuid.New()
	return
}
