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
	ID   uint      `gorm:"primaryKey,autoIncrement" json:"-"`
	UUID uuid.UUID `gorm:"type:varchar(60);" json:"uuid"`

	Name     string `gorm:"type:varchar(255)" json:"name"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"-"`

	Birthday *sql.NullTime   `json:"birthday"`
	Bio      *sql.NullString `json:"bio"`
	Link     *sql.NullString `json:"link"`

	Role string `gorm:"type:varchar(50);default:user"`

	ActivatedAt sql.NullTime `json:"-"`
	VerifiedAt  sql.NullTime `json:"-"`

	Meta      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Users struct
type Users struct {
	Users []User `json:"users"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.UUID = uuid.New()
	return
}
