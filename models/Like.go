/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://afaan.dev

*/

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Like struct {
	ID   uint      `gorm:"primaryKey,autoIncrement" json:"-"`
	UUID uuid.UUID `gorm:"type:varchar(60); uniqueIndex" json:"uuid"`

	UserId uint `gorm:"index" json:"-"`
	PostId uint `gorm:"index" json:"-"`

	Meta      string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (like *Like) BeforeCreate(tx *gorm.DB) (err error) {
	like.UUID = uuid.New()
	return
}
