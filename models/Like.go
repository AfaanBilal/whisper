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
	ID     uint      `gorm:"primaryKey,autoIncrement"`
	UUID   uuid.UUID `gorm:"type:varchar(60);"`
	UserId uint
	PostId uint

	Meta string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Likes struct
type Likes struct {
	Likes []Like `json:"likes"`
}

func (like *Like) BeforeCreate(tx *gorm.DB) (err error) {
	like.UUID = uuid.New()
	return
}
