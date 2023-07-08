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

type Post struct {
	ID        uint      `gorm:"primaryKey,autoIncrement"`
	UUID      uuid.UUID `gorm:"type:varchar(60);"`
	UserId    uint
	ReplyToId sql.NullInt64

	Content string
	Media   sql.NullString

	Meta string

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Posts struct
type Posts struct {
	Posts []Post `json:"posts"`
}

func (post *Post) BeforeCreate(tx *gorm.DB) (err error) {
	post.UUID = uuid.New()
	return
}
