/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID   uint      `gorm:"primaryKey,autoIncrement" json:"-"`
	UUID uuid.UUID `gorm:"type:varchar(60); uniqueIndex" json:"uuid"`

	Username string `gorm:"type:varchar(255); uniqueIndex" json:"username"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"type:varchar(255); uniqueIndex" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"-"`

	Image    string       `json:"image"`
	Birthday sql.NullTime `json:"birthday"`
	Bio      string       `json:"bio"`
	Link     string       `json:"link"`

	IsPrivate bool   `gorm:"index" json:"is_private"`
	Role      string `gorm:"type:varchar(50); default:user"`

	ActivatedAt sql.NullTime `gorm:"index" json:"activated_at"`
	VerifiedAt  sql.NullTime `gorm:"index" json:"verified_at"`

	Meta      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.UUID = uuid.New()
	return
}
