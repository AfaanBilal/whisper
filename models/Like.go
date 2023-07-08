/*

Whisper

A micro-blogging platform.

@author    Afaan Bilal
@copyright 2023 Afaan Bilal
@link      https://eonyx.io

*/

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Like struct {
	ID     uint      `gorm:"primaryKey,autoIncrement" json:"-"`
	UUID   uuid.UUID `gorm:"type:varchar(60);" json:"uuid"`
	UserId uint      `json:"-"`
	PostId uint      `json:"-"`

	Meta      *string   `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Likes struct
type Likes struct {
	Likes []Like `json:"likes"`
}

func (like *Like) BeforeCreate(tx *gorm.DB) (err error) {
	like.UUID = uuid.New()
	return
}
